package api

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gostack-labs/adminx/internal/code"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/adminx/internal/verifycode"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgconn"
)

type SignupRequest struct {
	Username   string `json:"username" validate:"required,alphanum"`
	Password   string `json:"password" validate:"required,min=6"`
	FullName   string `json:"full_name" validate:"required"`
	Email      string `json:"email" validate:"required_without=Phone,omitempty,email"`
	Phone      string `json:"phone" validate:"required_without=Email,omitempty,phone"`
	VerifyCode string `json:"verify_code" validate:"required,alphanum"`
}

type userResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}

func newUserResponse(user *db.User) userResponse {
	return userResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		Phone:            user.Phone,
		PasswordChangeAt: user.PasswordChangeAt,
		CreatedAt:        user.CreatedAt,
	}
}

func (server *Server) signup(c *bytego.Ctx) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	if strings.TrimSpace(req.Email) != "" {
		if ok := verifycode.NewVerifyCode().CheckAnswer(req.Email, req.VerifyCode); !ok {
			return resp.Fail(http.StatusBadRequest, code.VerifyCodeError).JSON(c)
		}
		b, err := server.store.CheckUserEmail(c.Context(), req.Email)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}
		if b {
			return resp.Fail(http.StatusFound, code.UserEmailExistError).JSON(c)
		}
	}

	if strings.TrimSpace(req.Phone) != "" {
		b, err := server.store.CheckUserPhone(c.Context(), req.Phone)
		if err != nil {
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
		}

		if b {
			return resp.Fail(http.StatusFound, code.UserPhoneExistError).JSON(c)
		}
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		Phone:          req.Phone,
	}

	user, err := server.store.CreateUser(c.Context(), arg)
	if err != nil {
		var pgxerr *pgconn.PgError
		if errors.As(err, &pgxerr) {
			if pgxerr.Code == "23505" {
				return resp.Fail(http.StatusFound, code.UserUsernameExistError).JSON(c)
			}
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}
	rsp := newUserResponse(user)
	return resp.OperateOK(rsp).JSON(c)
}

type VerifyCodeEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (s *Server) sendUsingEmail(c *bytego.Ctx) error {
	var req VerifyCodeEmailRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	err := verifycode.NewVerifyCode().SendEmail(req.Email)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.SendEmailError).WithError(err).JSON(c)
	}
	return resp.OperateOK().JSON(c)
}

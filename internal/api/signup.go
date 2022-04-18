package api

import (
	"database/sql"
	"net/http"
	"strings"
	"time"

	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/adminx/internal/verifycode"
	"github.com/gostack-labs/bytego"
	"github.com/lib/pq"
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

func newUserResponse(user db.User) userResponse {
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
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if strings.TrimSpace(req.Email) != "" {
		u, err := server.store.GetUserByEmail(c.Context(), req.Email)
		if err != nil {
			if err != sql.ErrNoRows {
				return c.JSON(http.StatusInternalServerError, errorResponse(err))
			}
		}
		if u != (db.User{}) {
			return c.JSON(http.StatusForbidden, bytego.Map{"error": "该邮箱已存在！"})
		}
	}

	if strings.TrimSpace(req.Phone) != "" {
		u, err := server.store.GetUserByPhone(c.Context(), req.Phone)
		if err != nil {
			if err != sql.ErrNoRows {
				return c.JSON(http.StatusInternalServerError, errorResponse(err))
			}
		}

		if u != (db.User{}) {
			return c.JSON(http.StatusForbidden, bytego.Map{"error": "该手机号已存在！"})
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
	}
	rsp := newUserResponse(user)
	return c.JSON(http.StatusOK, rsp)
}

type VerifyCodeEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (s *Server) sendUsingEmail(c *bytego.Ctx) error {
	var req VerifyCodeEmailRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	err := verifycode.NewVerifyCode().SendEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return c.JSON(http.StatusOK, bytego.Map{
		"success": true,
		"message": "发送成功",
	})
}

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

type signupRequest struct {
	Username   string `json:"username" validate:"required,alphanum"`                   // 用户名
	Password   string `json:"password" validate:"required,min=6"`                      // 密码
	FullName   string `json:"full_name" validate:"required"`                           // 全名
	Email      string `json:"email" validate:"required_without=Phone,omitempty,email"` // 邮箱
	Phone      string `json:"phone" validate:"required_without=Email,omitempty,phone"` // 手机号
	VerifyCode string `json:"verify_code" validate:"required,alphanum"`                // 验证码
} // 用户注册请求参数

type userResponse struct {
	Username         string    `json:"username"`           // 用户名
	FullName         string    `json:"full_name"`          // 全名
	Email            string    `json:"email"`              // 邮箱
	Phone            string    `json:"phone"`              // 手机号
	PasswordChangeAt time.Time `json:"password_change_at"` // 密码修改时间
	CreatedAt        time.Time `json:"created_at"`         // 创建时间
} // 用户注册返回数据

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

//@title 用户注册接口
//@api post /signup
//@group basic
//@request signupRequest
//@response 200 resp.resultOK{businesscode=10000,message="操作成功",data=userResponse}
func (server *Server) signup(c *bytego.Ctx) error {
	var req signupRequest
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

type verifyCodeEmailRequest struct {
	Email string `json:"email" validate:"required,email"` // 邮箱
} // 发送邮箱验证码请求参数

//@title 发送邮箱验证码接口
//@api get /signup/sendUsingEmail
//@group basic
//@request verifyCodeEmailRequest
//@response 200 resp.resultOK{businesscode=10000,message="操作成功"}
func (s *Server) sendUsingEmail(c *bytego.Ctx) error {
	var req verifyCodeEmailRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	err := verifycode.NewVerifyCode().SendEmail(req.Email)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.SendEmailError).WithError(err).JSON(c)
	}
	return resp.OperateOK().JSON(c)
}

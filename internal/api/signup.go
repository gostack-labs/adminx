package api

import (
	"net/http"
	"time"

	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/bytego"
	"github.com/lib/pq"
)

type SignupByEmailRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type SignupResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}

func newSignupResponse(user db.User) SignupResponse {
	return SignupResponse{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		Phone:            user.Phone,
		PasswordChangeAt: user.PasswordChangeAt,
		CreatedAt:        user.CreatedAt,
	}
}

func (server *Server) signupByEmail(c *bytego.Ctx) error {
	var req SignupByEmailRequest
	if err := c.Bind(&req); err != nil {
		_ = c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	// TODO email unique check

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(c.Request.Context(), arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				_ = c.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
	}
	rsp := newSignupResponse(user)
	_ = c.JSON(http.StatusOK, rsp)
	return err
}

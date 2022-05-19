package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/code"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/bytego"
	"github.com/jackc/pgx/v4"
)

type logginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type logginUserResponse struct {
	SessinID              uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) logginUser(c *bytego.Ctx) error {
	var req logginUserRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}
	user, err := server.store.GetUser(c.Context(), req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return resp.Fail(http.StatusNotFound, code.UserNotExistError).JSON(c)
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return resp.Fail(http.StatusForbidden, code.UserPwdError).JSON(c)
	}

	accessToken, accessPaylod, err := server.tokenMaker.CreateToken(
		user.Username,
		configs.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.TokenCreateError).WithError(err).JSON(c)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		configs.Get().Token.RefreshTokenDuration,
	)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.RefreshTokenCreateError).WithError(err).JSON(c)
	}

	session, err := server.store.CreateSession(c.Context(), db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    c.Request.UserAgent(),
		ClientIp:     c.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	rsp := logginUserResponse{
		SessinID:              session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPaylod.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	return resp.OperateOK(rsp).JSON(c)
}

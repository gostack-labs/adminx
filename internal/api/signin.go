package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gostack-labs/adminx/configs"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/utils"
	"github.com/gostack-labs/bytego"
)

type logginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
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
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := server.store.GetUser(c, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, errorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	accessToken, accessPaylod, err := server.tokenMaker.CreateToken(
		user.Username,
		configs.Config.Token.AccessTokenDuration,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		configs.Config.Token.RefreshTokenDuration,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	session, err := server.store.CreateSession(c, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    c.Request.UserAgent(),
		ClientIp:     c.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := logginUserResponse{
		SessinID:              session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPaylod.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	return c.JSON(http.StatusOK, rsp)
}

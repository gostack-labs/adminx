package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/bytego"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(c *bytego.Ctx) error {
	var req renewAccessTokenRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	session, err := server.store.GetSession(c, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, errorResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	if session.Username != refreshPayload.Username {
		err := fmt.Errorf("incorrect session user")
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return c.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		configs.Config.Token.AccessTokenDuration,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	return c.JSON(http.StatusOK, rsp)
}

package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/adminx/pkg/token"
	"github.com/gostack-labs/bytego"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(c *bytego.Ctx) error {
	var req renewAccessTokenRequest
	if err := c.Bind(&req); err != nil {
		return resp.BadRequestJSON(err, c)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, token.ErrExpiredToken) {
			return resp.Fail(http.StatusUnauthorized, code.TokenExpiredError).JSON(c)
		}
		if errors.Is(err, token.ErrInvalidToken) {
			return resp.Fail(http.StatusUnauthorized, code.TokenInvalidError).JSON(c)
		}
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	session, err := server.store.GetSession(c.Context(), refreshPayload.ID)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	if session.IsBlocked {
		return resp.Fail(http.StatusUnauthorized, code.SessionBlockedError).JSON(c)
	}

	if session.Username != refreshPayload.Username {
		return resp.Fail(http.StatusUnauthorized, code.SessionError).JSON(c)
	}

	if session.RefreshToken != req.RefreshToken {
		return resp.Fail(http.StatusUnauthorized, code.SessionMismatchedError).JSON(c)
	}

	if time.Now().After(session.ExpiresAt) {
		return resp.Fail(http.StatusUnauthorized, code.SessionExpiredError).JSON(c)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		configs.Get().Token.AccessTokenDuration,
	)
	if err != nil {
		return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).JSON(c)
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	return resp.OperateOK(rsp).JSON(c)
}

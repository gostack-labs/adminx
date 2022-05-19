package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gostack-labs/adminx/internal/code"
	"github.com/gostack-labs/adminx/internal/resp"
	"github.com/gostack-labs/adminx/pkg/token"
	"github.com/gostack-labs/bytego"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) bytego.HandlerFunc {
	return func(c *bytego.Ctx) error {
		authorizationHeader := c.Header(AuthorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			return resp.Fail(http.StatusUnauthorized, code.AuthHeaderNotExistError).AbortJSON(c)
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			return resp.Fail(http.StatusUnauthorized, code.AuthHeaderFormatError).AbortJSON(c)
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			return resp.Fail(http.StatusUnauthorized, code.AuthTypeError).AbortJSON(c)
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			if errors.Is(err, token.ErrExpiredToken) {
				return resp.Fail(http.StatusUnauthorized, code.TokenExpiredError).AbortJSON(c)
			}
			if errors.Is(err, token.ErrInvalidToken) {
				return resp.Fail(http.StatusUnauthorized, code.TokenInvalidError).AbortJSON(c)
			}
			return resp.Fail(http.StatusInternalServerError, code.ServerError).WithError(err).AbortJSON(c)
		}

		c.Set(AuthorizationPayloadKey, payload)
		return c.Next()
	}
}

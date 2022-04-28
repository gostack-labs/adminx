package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gostack-labs/adminx/pkg/common"
	"github.com/gostack-labs/adminx/pkg/token"
	"github.com/gostack-labs/bytego"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) bytego.HandlerFunc {
	return func(c *bytego.Ctx) error {
		authorizationHeader := c.Header(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return common.NewCommonError(http.StatusUnauthorized, "authorization header is not provided")
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return common.NewCommonError(http.StatusUnauthorized, "invalid authorization header format")
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			c.AbortWithStatus(http.StatusUnauthorized)
			return common.NewCommonError(http.StatusUnauthorized, fmt.Sprintf("unsupported authorization type %s", authorizationType))
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.Abort()
			return err
		}

		c.Set(AuthorizationPayloadKey, payload)
		return c.Next()
	}
}

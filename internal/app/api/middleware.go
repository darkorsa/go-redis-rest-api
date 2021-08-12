package api

import (
	"fmt"
	"strings"

	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/darkorsa/go-redis-http-client/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := apiErrors.NewUnauthorizedError("authorization header is not provided")
			c.AbortWithStatusJSON(err.GetStatus(), err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := apiErrors.NewUnauthorizedError("invalid authorization header format")
			c.AbortWithStatusJSON(err.GetStatus(), err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := apiErrors.NewUnauthorizedError(fmt.Sprintf("unsupported authorization type %s", authorizationType))
			c.AbortWithStatusJSON(err.GetStatus(), err)
			return
		}

		accessToken := fields[1]
		payload, verErr := tokenMaker.VerifyToken(accessToken)
		if verErr != nil {
			err := apiErrors.NewUnauthorizedError(verErr.Error())
			c.AbortWithStatusJSON(err.GetStatus(), err)
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}

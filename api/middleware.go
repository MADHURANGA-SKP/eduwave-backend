package api

import (
	"eduwave-back-end/token"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func authMiddleware(tokenMaker token.Maker, accessbileRoles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if authorization header is provided
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		// Check if authorization header is empty
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)

		// Check if authorization header has the valid format or not
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])

		// Check if authorization type is supported
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)

		// Handle verification errors
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		if !hasPermission(payload.Role, accessbileRoles){
			err := fmt.Errorf("permision denied")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func hasPermission(userRole string, accessbileRoles []string) bool {
	for _, role := range accessbileRoles{
		if userRole == role {
			return true
		}
	}
	return false
}


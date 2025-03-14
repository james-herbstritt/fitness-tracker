package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}

		// Expect "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}

		accessToken := tokenParts[1]

		if IsTokenExpired(accessToken) {
			// get a new token
			newToken, err := RefreshToken(accessToken)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired and could not be refreshed"})
				return
			}

			// Set new token in context
			accessToken = newToken
			ctx.Header("Authorization", "Bearer "+newToken)
		}

		// Pass valid token to handlers
		ctx.Set("accessToken", accessToken)
		ctx.Next()
	}
}

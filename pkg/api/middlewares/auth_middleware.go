package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/env"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/jwt"
)

// Function to ValidateToken takes the jwt token from the request header and checks the validity of the token
func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the jwt token from the request header
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Get the jwt secret from the environment variable
		jwtSecret := env.EnvConfig.JWTSecretKey
		if jwtSecret == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "jwt secret not set"})
			ctx.Abort()
			return
		}

		// Validate the token
		err := jwt.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

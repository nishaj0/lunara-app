package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/jwt"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"go.uber.org/zap"
)

// AuthMiddleware validates JWT tokens and sets user context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Missing authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check if header starts with "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			logger.Warn("Invalid authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := tokenParts[1]

		// Validate token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			logger.Warn("Invalid token", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)

		logger.Info("User authenticated", zap.String("user_id", claims.UserID), zap.String("email", claims.Email))

		c.Next()
	}
}

// GetUserFromContext extracts user information from gin context
func GetUserFromContext(c *gin.Context) (userID, email string, exists bool) {
	userIDInterface, exists1 := c.Get("user_id")
	emailInterface, exists2 := c.Get("email")

	if !exists1 || !exists2 {
		return "", "", false
	}

	userID, ok1 := userIDInterface.(string)
	email, ok2 := emailInterface.(string)

	if !ok1 || !ok2 {
		return "", "", false
	}

	return userID, email, true
}
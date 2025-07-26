package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/middleware"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"go.uber.org/zap"
)

// GetProfile returns the current user's profile information
func GetProfile(c *gin.Context) {
	// get user information from context (set by AuthMiddleware)
	userID, email, exists := middleware.GetUserFromContext(c)
	if !exists {
		logger.Error("User context not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User context not found"})
		return
	}

	logger.Info("Profile requested", zap.String("user_id", userID), zap.String("email", email))

	// send user profile data, for now just returning user ID and email
	// after creating user page, this can be expanded to include more user details that are fetched from the database
	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"email":   email,
		"message": "Profile data retrieved successfully",
	})
}
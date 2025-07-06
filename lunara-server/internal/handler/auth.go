package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/model"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"github.com/nishaj0/lunara-app/lunara-server/internal/service"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid register request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Register attempt", zap.String("email", req.Email), zap.String("username", req.Username))

	user, err := service.RegisterUser(c.Request.Context(), &req)
	if err != nil {
		logger.Error("Failed to register user", zap.Error(err), zap.String("email", req.Email))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("User registered successfully", zap.String("user_id", user.ID), zap.String("email", user.Email))

	c.JSON(http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"fullName": user.FullName,
	})
}

func Login(c *gin.Context) {
	var req model.LoginRequest
	environment := env.GetEnv("ENV", "DEV")
	
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Warn("Invalid login request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := service.LoginUser(c.Request.Context(), &req)
	if err != nil {
		logger.Warn("Login failed", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Set JWT as a secure, HTTP-only cookie
	c.SetCookie(
		"token",                     // cookie name
		token,                       // value
		60*60*24,                    // maxAge (1 day in seconds)
		"/",                         // path
		"",                          // domain (empty = current domain)
		environment == "PROD",       // secure (set to true in production)
		true,                        // httpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"fullName": user.FullName,
	})
}

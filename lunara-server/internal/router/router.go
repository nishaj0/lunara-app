package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/handler"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"github.com/nishaj0/lunara-app/lunara-server/internal/middleware"
)

func SetupRouter() *gin.Engine {
	// initialize logger
	logger.InitLogger()

	// set Gin mode
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Recovery middleware
	r.Use(gin.Recovery())
	
	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	// custom logger middleware
	r.Use(middleware.Logger())
	
	// public routes
	r.GET("/ping", handler.Ping)
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)

	// protected routes (require authentication)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handler.GetProfile)
		// Add more protected routes here as needed
	}

	return r
}
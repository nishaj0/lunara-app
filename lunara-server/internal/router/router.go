package router

import (
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
	
	// custom logger middleware
	r.Use(middleware.Logger())
	
	// general routes
	r.GET("/ping", handler.Ping)
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)

	return r
}
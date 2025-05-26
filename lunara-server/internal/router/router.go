package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// general routes
	r.GET("/ping", handler.Ping)

	return r
}
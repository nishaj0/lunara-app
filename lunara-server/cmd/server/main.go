package main

import (
	"fmt"

	"github.com/nishaj0/lunara-app/lunara-server/internal/db"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"github.com/nishaj0/lunara-app/lunara-server/internal/router"
	"go.uber.org/zap"
)

func main() {
	// Load environment variables first
	env.LoadEnv()

	
	host := env.GetEnv("HOST", "localhost")
	port := env.GetEnv("PORT", "8080")

	// Initialize database
	if err := db.InitDB(); err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
	}
	defer db.CloseDB()

	r := router.SetupRouter()

	logger.Info("Server is starting", zap.String("host", host), zap.String("port", port))

	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
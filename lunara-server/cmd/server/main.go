package main

import (
	"fmt"
	"log"

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

	log.Printf("Server is starting at %s:%s", host, port)

	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
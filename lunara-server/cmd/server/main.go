package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nishaj0/lunara-app/lunara-server/internal/router"
)

func main() {
	host := getEnv("HOST", "localhost")
	port := getEnv("PORT", "8080")

	r := router.SetupRouter()

	log.Printf("Server is starting at %s:%s", host, port)

	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
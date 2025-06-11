package main

import (
	"fmt"
	"log"

	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
	"github.com/nishaj0/lunara-app/lunara-server/internal/router"
)

func main() {
	host := env.GetEnv("HOST", "localhost")
	port := env.GetEnv("PORT", "8080")

	r := router.SetupRouter()

	log.Printf("Server is starting at %s:%s", host, port)

	err := r.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
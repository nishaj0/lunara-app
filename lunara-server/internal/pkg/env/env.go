package env

import (
	"os"

	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"go.uber.org/zap"
	"github.com/joho/godotenv"
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        logger.Warn("Warning: .env file not found or error loading it", zap.Error(err))
    }
}

// GetEnv retrieves an environment variable value with a fallback default
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != ""	{
		return value
	}
	return fallback
}
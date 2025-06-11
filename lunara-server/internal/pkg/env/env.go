package env

import "os"

// GetEnv retrieves an environment variable value with a fallback default
func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != ""	{
		return value
	}
	return fallback
}
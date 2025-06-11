package config

import (
	"fmt"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/env"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// NewDatabaseConfig initializes a new DatabaseConfig with values from environment variables.
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     env.GetEnv("DB_HOST", "localhost"),
		Port:     env.GetEnv("DB_PORT", "5432"),
		User:     env.GetEnv("DB_USER", "postgres"),
		Password: env.GetEnv("DB_PASSWORD", ""),
		DBName:   env.GetEnv("DB_NAME", "lunara"),
	}
}

// GetDSN constructs the Data Source Name (DSN) for connecting to the database.
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

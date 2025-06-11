package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nishaj0/lunara-app/lunara-server/internal/config"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"go.uber.org/zap"
)

var dbPool *pgxpool.Pool

func InitDB() error {
	dbConfig := config.NewDatabaseConfig()
	poolConfig, err := pgxpool.ParseConfig(dbConfig.GetDSN())
	if err != nil {
		return err
	}

	// Set pool configuration
	poolConfig.MaxConns = 10
	poolConfig.MinConns = 2

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return err
	}

	// Verify connection
	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}

	dbPool = pool
	logger.Info("Database connection established", zap.String("dsn", dbConfig.GetDSN()))
	return nil
}

func GetDB() *pgxpool.Pool {
	return dbPool
}

func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
		logger.Info("Database connection closed")
	}
}

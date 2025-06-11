package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nishaj0/lunara-app/lunara-server/internal/pkg/logger"
	"go.uber.org/zap"
)

// logger middleware for Gin
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// process request
		c.Next()

		// log request details
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		logger.GetLogger().Info("HTTP Request",
			zap.String("path", path),
			zap.String("query", query),
			zap.String("method", method),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		)
	}
}

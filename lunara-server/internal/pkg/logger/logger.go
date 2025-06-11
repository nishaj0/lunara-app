package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

// InitLogger initializes the logger
func InitLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	
	// Customize the logging configuration
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	
	// Set the log level 
	logLevel := "info"
	if logLevel != "" {
		level, err := zapcore.ParseLevel(logLevel)
		if err == nil {
			config.Level.SetLevel(level)
		}
	}

	var err error
	log, err = config.Build()
	if err != nil {
		panic(err)
	}

	return log
}

// GetLogger returns the logger instance
func GetLogger() *zap.Logger {
	if log == nil {
		log = InitLogger()
	}
	return log
}

// Info logs an info message
func Info(message string, fields ...zapcore.Field) {
	GetLogger().Info(message, fields...)
}

// Warn logs a warning message
func Warn(message string, fields ...zapcore.Field) {
	GetLogger().Warn(message, fields...)
}

// Error logs an error message
func Error(message string, fields ...zapcore.Field) {
	GetLogger().Error(message, fields...)
}

// Debug logs a debug message
func Debug(message string, fields ...zapcore.Field) {
	GetLogger().Debug(message, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(message string, fields ...zapcore.Field) {
	GetLogger().Fatal(message, fields...)
}

// With creates a child logger with additional fields
func With(fields ...zapcore.Field) *zap.Logger {
	return GetLogger().With(fields...)
}
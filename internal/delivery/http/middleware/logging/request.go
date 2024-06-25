package logging

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/pkg/logging"
	"log/slog"
	"time"
)

func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log := logging.MustGetLogger(c)

		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate response time
		duration := time.Since(startTime).Milliseconds()

		// Extract request attributes
		logEntry := log.With(
			slog.String("action", "request_processed"),
			slog.String("client_ip", c.ClientIP()),
			slog.Int64("response_time", duration),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.RequestURI),
			slog.Int("response_status", c.Writer.Status()),
		)

		// Log the information in JSON format
		logEntry.Info("Request Processed")
	}
}

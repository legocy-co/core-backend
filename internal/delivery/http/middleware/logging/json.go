package logging

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func JSONLogMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate response time
		duration := time.Since(startTime).Milliseconds()

		// Extract request attributes
		logEntry := log.With(
			slog.String("client_ip", c.ClientIP()),
			slog.Int64("response_time", duration),
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.RequestURI),
			slog.String("request_id", c.Writer.Header().Get("Request-Id")),
			slog.Int("response_status", c.Writer.Status()),
		)

		// Log the information in JSON format
		logEntry.Info("Request Processed")
	}
}

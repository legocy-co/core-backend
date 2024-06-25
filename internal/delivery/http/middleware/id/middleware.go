package id

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/legocy-co/legocy/internal/pkg/logging"
	"log/slog"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()

		// Set Header
		c.Request.Header.Set("X-Request-ID", requestID)

		// Modify logger
		log := logging.MustGetLogger(c)
		log = log.With(slog.String("request_id", requestID))
		c.Set("log", log)

		c.Next()
	}
}

package id

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/legocy-co/legocy/internal/pkg/logging"
	"log/slog"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("X-Request-ID", uuid.New().String())

		log := logging.MustGetLogger(c)
		log = log.With(slog.String("request_id", c.GetHeader("X-Request-ID")))

		c.Next()
	}
}

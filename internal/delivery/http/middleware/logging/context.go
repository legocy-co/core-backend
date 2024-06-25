package logging

import (
	"github.com/gin-gonic/gin"
	"log/slog"
)

func ContextLoggerMiddleware(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("log", log)
		c.Next()
	}
}

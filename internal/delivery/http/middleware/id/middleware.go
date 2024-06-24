package id

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("X-Request-ID", uuid.New().String())
		c.Next()
	}
}

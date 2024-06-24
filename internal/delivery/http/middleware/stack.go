package middleware

import "github.com/gin-gonic/gin"

func NewStack(funcs ...gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, f := range funcs {
			if f != nil {
				f(c)
			}
		}
	}
}

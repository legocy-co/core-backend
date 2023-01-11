package v1

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("templates", "templates")

	return r
}

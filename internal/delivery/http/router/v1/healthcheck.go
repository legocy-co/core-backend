package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r V1router) addHealthCheck(rg *gin.RouterGroup) {

	healthcheck := rg.Group("/health")
	{
		healthcheck.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}
}

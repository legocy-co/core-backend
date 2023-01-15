package v1

import (
	h "legocy-go/api/v1/handlers"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouter(userService s.UserUseCase, seriesService s.LegoSeriesService) *gin.Engine {
	r := gin.Default()

	tokenHandler := h.NewTokenHandler(userService)
	legoSeriesHandler := h.NewLegoSeriesHandler(seriesService)

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/token", tokenHandler.GenerateToken)
			auth.POST("/register", tokenHandler.UserRegister)
		}
		series := v1.Group("/series").Use(m.Auth())
		{
			series.GET("/", legoSeriesHandler.ListSeries)
			series.POST("/", legoSeriesHandler.SeriesCreate)
			series.GET("/:seriesID", legoSeriesHandler.DetailSeries)
			series.DELETE("/:seriesID", legoSeriesHandler.DeleteSeries)
		}
	}

	return r
}

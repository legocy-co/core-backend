package v1

import (
	s "legocy-go/api/v1/usecase/auth"
	"legocy-go/api/v1/usecase/lego"

	"github.com/gin-gonic/gin"
)

type V1router struct {
	router *gin.Engine
}

func (r V1router) Run(port string) error {
	return r.router.Run(":" + port)
}

func InitRouter(
	userService s.UserUseCase,
	legoSeriesService lego.LegoSeriesService,
	legoSetService lego.LegoSetUseCase) V1router {

	r := gin.Default()
	router := V1router{router: r}

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

	//legoseries.go
	router.addLegoSeries(v1, legoSeriesService)
	//auth.go
	router.addAuth(v1, userService)

	router.addLegoSets(v1, legoSetService)

	return router
}

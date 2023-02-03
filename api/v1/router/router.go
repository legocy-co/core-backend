package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/app"
)

type V1router struct {
	router *gin.Engine
}

func (r V1router) Run(port string) error {
	return r.router.Run(":" + port)
}

func InitRouter(app *app.App) V1router {

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

	//auth.go
	router.addAuth(v1, app.GetUserService())

	//user_images.go
	router.addUserImages(v1, app.GetUserImagesService(), app.GetStorage())

	//legoseries.go
	router.addLegoSeries(v1, app.GetLegoSeriesService())

	//legoset.go
	router.addLegoSets(v1, app.GetLegoSetService())

	//location.go
	router.addLocations(v1, app.GetLocationService())

	//currency.go
	router.addCurrencies(v1, app.GetCurrencyService())

	//marketitem.go
	router.addMarketItems(v1, app)

	return router
}

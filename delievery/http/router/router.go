package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "legocy-go/docs"
	"legocy-go/internal/app"
)

type V1router struct {
	router *gin.Engine
}

func (r V1router) Run(port string) error {
	return r.router.Run(":" + port)
}

//	@title			LEGOcy API
//	@version		1.0
//	@description	LEGOcy is a marketplace for LEGO lovers.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host		localhost:8080
//	@BasePath	/api/v1
//
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
//
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

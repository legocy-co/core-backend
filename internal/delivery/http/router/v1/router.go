package v1

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/app"
	"github.com/legocy-co/legocy/internal/delivery/http/router/v1/admin"
	"github.com/legocy-co/legocy/internal/delivery/http/router/v1/site"
	"github.com/legocy-co/legocy/pkg/logging"
	"github.com/legocy-co/legocy/pkg/logging/util"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
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

// GetV1Router
// @contact.name				API Support
// @contact.url				http://www.swagger.io/support
// @contact.email				support@swagger.io
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath					/api/v1
//
// @securityDefinitions.apiKey	JWT
// @in							header
// @name						Authorization
//
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func GetV1Router(app *app.App) V1router {

	gin.SetMode(gin.ReleaseMode)
	util.UseJSONLogFormat()

	r := gin.Default()

	r.Use(logging.JSONLogMiddleware())

	router := V1router{router: r}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1Admin := r.Group("/api/v1/admin")

	// Admin Routers
	admin.AddUserAdmin(v1Admin, app.GetUserAdminService())
	admin.AddAdminMarketItems(v1Admin, app.GetMarketItemAdminService())
	admin.AddAdminLegoSetValuations(v1Admin, app.GetLegoSetValuationAdminService())

	// Site Routers
	site.AddImagesRoutes(v1, app)
	site.AddUsers(v1, app)
	site.AddUserCollections(v1, app)
	site.AddLegoSeries(v1, app.GetLegoSeriesService())
	site.AddLegoSets(v1, app.GetLegoSetService())
	site.AddMarketItems(v1, app)

	// healthcheck.go
	router.addHealthCheck(v1)

	return router
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/routes/admin"
	"github.com/legocy-co/legocy/internal/delivery/http/routes/site"
	"github.com/legocy-co/legocy/internal/delivery/http/routes/swagger"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func AddRoutes(engine *gin.Engine, app *app.App) {
	addAdminRoutes(engine, app)
	addSiteRoutes(engine, app)
	addSwaggerRoutes(engine, app)
}

func addAdminRoutes(engine *gin.Engine, app *app.App) {
	v1Admin := engine.Group("/api/v1/admin")

	admin.AddUserAdmin(v1Admin, app.GetUserAdminService())
	admin.AddAdminMarketItems(v1Admin, app)
	admin.AddAdminLegoSetValuations(v1Admin, app.GetLegoSetValuationAdminService())
}

func addSiteRoutes(engine *gin.Engine, app *app.App) {
	v1 := engine.Group("/api/v1")

	site.AddUsers(v1, app)
	site.AddUserCollections(v1, app)
	site.AddLegoSeries(v1, app.GetLegoSeriesService())
	site.AddLegoSets(v1, app)
	site.AddMarketItems(v1, app)
	site.AddCallcuatorRoutes(v1, app)
}

func addSwaggerRoutes(engine *gin.Engine, app *app.App) {
	swagger.AddSwaggerDocs(engine)
}

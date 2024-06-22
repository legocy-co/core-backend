package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/legocy-co/legocy/docs"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/id"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/logging"
	"github.com/legocy-co/legocy/internal/delivery/http/routes"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

type Server struct {
	engine *gin.Engine
}

//	@title			LEGOcy API
//	@version		1.0
//	@description	LEGOcy is a marketplace for LEGO lovers.
//	@termsOfService	http://swagger.io/terms/

// New
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
func New(app *app.App) *Server {

	e := gin.Default()

	// Use middleware
	e.Use(
		id.RequestIDMiddleware(),
		logging.JSONLogMiddleware(app.GetLogger()),
	)

	// Add routes
	routes.AddRoutes(e, app)

	return &Server{engine: e}
}

func (s *Server) Run(port string) error {
	return s.engine.Run(":" + port)
}

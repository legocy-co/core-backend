package site

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/calculator"
	"github.com/legocy-co/legocy/internal/pkg/app"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddCallcuatorRoutes(rg *gin.RouterGroup, app *app.App) {
	handler := calculator.NewLegoSetValuationHandler(app.GetLegoSetValuationService())

	calculatorRoutes := rg.Group("/sets-valuations").Use(middleware.IsAuthenticated())
	{
		calculatorRoutes.GET("/:legoSetID", handler.GetValuationsByLegoSetID)
	}
}

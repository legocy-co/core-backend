package site

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/calculator"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func AddCallcuatorRoutes(rg *gin.RouterGroup, app *app.App) {
	handler := calculator.NewLegoSetValuationHandler(app.GetLegoSetValuationService())

	calculatorRoutes := rg.Group("/sets-valuations").Use(auth.IsAuthenticated())
	{
		calculatorRoutes.GET("/:legoSetID", handler.GetValuationsByLegoSetID)
	}
}

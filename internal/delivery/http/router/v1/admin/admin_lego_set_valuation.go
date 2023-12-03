package admin

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/calculator/admin"
	"github.com/legocy-co/legocy/internal/domain/calculator/service/admin"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddAdminLegoSetValuations(rg *gin.RouterGroup, service admin.LegoSetValuationAdminService) {
	handler := h.NewHandler(service)

	valuationsAdmin := rg.Group("/sets-valuations").Use(middleware.IsAdmin())
	{
		valuationsAdmin.GET("/", handler.GetLegoSetValuations)
		valuationsAdmin.GET("/:valuationID", handler.GetValuationByID)
		valuationsAdmin.POST("/", handler.CreateValuation)
		valuationsAdmin.PUT("/:valuationID", handler.UpdateValuation)
		valuationsAdmin.DELETE("/:valuationID", handler.DeleteValuation)
	}

}

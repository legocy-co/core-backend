package v1

import (
	"github.com/gin-gonic/gin"
	h "legocy-go/internal/delivery/http/handlers/marketplace/admin/marketItem"
	s "legocy-go/internal/domain/marketplace/service/admin"
	"legocy-go/pkg/auth/jwt/middleware"
)

func (r V1router) addAdminMarketItems(rg *gin.RouterGroup, service s.MarketItemAdminService) {
	handler := h.NewMarketItemAdminHandler(service)

	marketItemsAdmin := rg.Group("/market-items").Use(middleware.IsAdmin())
	{
		marketItemsAdmin.GET("/", handler.GetMarketItemsAdmin)
		marketItemsAdmin.GET("/:itemId", handler.GetMarketItemByID)
		marketItemsAdmin.POST("/", handler.CreateMarketItem)
		marketItemsAdmin.PUT("/:itemId", handler.UpdateMarketItemByID)
		marketItemsAdmin.DELETE("/:itemId", handler.DeleteMarketItemById)
	}
}

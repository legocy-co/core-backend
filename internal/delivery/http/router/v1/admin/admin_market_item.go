package admin

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/admin/marketItem"
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service/admin"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddAdminMarketItems(rg *gin.RouterGroup, service s.MarketItemAdminService) {
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

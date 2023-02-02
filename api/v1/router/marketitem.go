package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/handlers/marketplace"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/marketplace"
	r "legocy-go/pkg/marketplace/repository"
)

func (r V1router) addMarketItems(
	rg *gin.RouterGroup,
	service s.MarketItemService, repo r.MarketItemRepository) {

	handler := marketplace.NewMarketItemHandler(service)

	items := rg.Group("/market-items").Use(m.Auth())
	{
		items.GET("/", handler.ListMarketItems)
		items.POST("/", handler.CreateMarketItem)
		items.Use(m.ItemOwnerOrAdmin("itemId", repo))
		{
			items.DELETE("/:itemId", handler.DeleteMarketItem)
		}
	}

}

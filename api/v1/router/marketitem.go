package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/handlers/marketplace"
	m "legocy-go/api/v1/middleware"
	a "legocy-go/internal/app"
)

func (r V1router) addMarketItems(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := marketplace.NewMarketItemHandler(app.GetMarketItemService())

	items := rg.Group("/market-items").Use(m.Auth())
	{
		items.GET("/", handler.ListMarketItems)
		items.Use(
			m.HasFreeMarketItemsSlot(a.MaxItemsOwnedByUser, app.GetMarketItemRepo()))
		{
			items.POST("/", handler.CreateMarketItem)
		}

		items.Use(m.ItemOwnerOrAdmin("itemId", app.GetMarketItemRepo()))
		{
			items.DELETE("/:itemId", handler.DeleteMarketItem)
		}
	}

}

package v1

import (
	"github.com/gin-gonic/gin"
	a "legocy-go/internal/app"
	"legocy-go/internal/delievery/http/handlers/marketplace"
	"legocy-go/internal/delievery/http/middleware"
)

func (r V1router) addMarketItems(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := marketplace.NewMarketItemHandler(
		app.GetMarketItemService(), app.GetNotifyEventClient())

	items := rg.Group("/market-items").Use(v1.Auth())
	{
		items.GET("/", handler.ListMarketItems)
		items.GET("/authorized/", handler.ListMarketItemsAuthorized)
		items.GET("/:itemID", handler.MarketItemDetail)

		items.Use(
			v1.HasFreeMarketItemsSlot(a.MaxItemsOwnedByUser, app.GetMarketItemRepo()))
		{
			items.POST("/", handler.CreateMarketItem)
		}
		items.Use(v1.ItemOwnerOrAdmin("itemId", app.GetMarketItemRepo()))
		{
			items.DELETE("/:itemId", handler.DeleteMarketItem)
		}
		items.Use(v1.IsMarketItemOwner("itemID", app.GetMarketItemRepo()))
		{
			items.PUT("/:itemID", handler.UpdateMarketItemByID)
		}
	}
	itemsAdmin := rg.Group("/admin/market-items").Use(v1.AdminUserOnly())
	{
		itemsAdmin.PUT("/:itemID", handler.UpdateMarketItemByIDAdmin)
	}
}

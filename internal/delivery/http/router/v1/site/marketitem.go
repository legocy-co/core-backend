package site

import (
	"github.com/gin-gonic/gin"
	a "github.com/legocy-co/legocy/internal/app"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware"
	jwt "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddMarketItems(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := marketplace.NewMarketItemHandler(
		app.GetMarketItemService())

	items := rg.Group("/market-items").Use(jwt.IsAuthenticated())
	{
		items.GET("/", handler.ListMarketItems)
		items.GET("/authorized/", handler.ListMarketItemsAuthorized)
		items.GET("/:itemID", handler.MarketItemDetail)

		items.Use(
			middleware.HasFreeMarketItemsSlot(a.MaxItemsOwnedByUser, app.GetMarketItemRepo()))
		{
			items.POST("/", handler.CreateMarketItem)
		}
		items.Use(middleware.ItemOwnerOrAdmin("itemId", app.GetMarketItemRepo()))
		{
			items.DELETE("/:itemId", handler.DeleteMarketItem)
		}
		items.Use(middleware.IsMarketItemOwner("itemID", app.GetMarketItemRepo()))
		{
			items.PUT("/:itemID", handler.UpdateMarketItemByID)
		}
	}
}

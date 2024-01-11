package site

import (
	"github.com/gin-gonic/gin"
	a "github.com/legocy-co/legocy/internal/app"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/image"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/market_item"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware"
	jwt "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddMarketItems(
	rg *gin.RouterGroup,
	app *a.App) {

	handler := market_item.NewMarketItemHandler(
		app.GetMarketItemService())

	items := rg.Group("/market-items")
	{
		items.GET("/", handler.ListMarketItems)

		items.Use(jwt.IsAuthenticated())
		{
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

	itemImages := rg.Group("/market-items/images")
	{
		handler := image.NewHandler(app.GetMarketItemImageService(), app.GetImageStorageClient())

		itemImages.Use(middleware.IsMarketItemOwner("marketItemID", app.GetMarketItemRepo()))
		{
			itemImages.POST("/:marketItemID", handler.UploadImage)
		}
	}

}

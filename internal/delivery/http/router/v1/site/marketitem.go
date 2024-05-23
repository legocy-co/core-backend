package site

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/image"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/like"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/market_item"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware"
	"github.com/legocy-co/legocy/internal/pkg/app"
	jwt "github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddMarketItems(rg *gin.RouterGroup, a *app.App) {

	items := rg.Group("/market-items")
	{
		handler := market_item.NewMarketItemHandler(a.GetMarketItemService(), a.GetUserReviewService())

		items.GET("/", handler.ListMarketItems)

		items.Use(jwt.IsAuthenticated())
		{
			items.GET("/authorized/", handler.ListMarketItemsAuthorized)
			items.GET("/:itemID", handler.MarketItemDetail)
			items.GET("/favorites/", handler.GetFavorites)

			privateRoutes := items.Group("")
			privateRoutes.Use(middleware.ItemOwnerOrAdmin("itemId", a.GetMarketItemRepo()))
			{
				privateRoutes.DELETE("/:itemId", handler.DeleteMarketItem)
				privateRoutes.PUT("/:itemId", handler.UpdateMarketItemByID)
			}

			checkSlotsRoutes := items.Group("")
			checkSlotsRoutes.Use(
				middleware.HasFreeMarketItemsSlot(app.MaxItemsOwnedByUser, a.GetMarketItemRepo()))
			{
				checkSlotsRoutes.POST("/", handler.CreateMarketItem)
			}
		}
	}

	itemImages := rg.Group("/market-items/images")
	{
		handler := image.NewHandler(a.GetMarketItemImageService(), a.GetImageStorageClient())

		itemImages.Use(middleware.IsMarketItemOwner("marketItemID", a.GetMarketItemRepo()))
		{
			itemImages.POST("/:marketItemID", handler.UploadImage)
			itemImages.DELETE("/:imageId", handler.Delete)
			itemImages.PATCH("/:marketItemID/:imageID", handler.Update)
		}
	}

	likeRoutes := rg.Group("/market-items/likes")
	{
		handler := like.NewHandler(a.GetMarketItemLikeRepository())

		likeRoutes.Use(jwt.IsAuthenticated())
		{
			likeRoutes.GET("/", handler.GetLikedItems)
			likeRoutes.POST("/:marketItemID", handler.LikeMarketItem)
			likeRoutes.DELETE("/:marketItemID", handler.UnlikeMarketItem)
		}
	}

}

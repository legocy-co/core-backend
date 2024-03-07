package admin

import (
	"github.com/gin-gonic/gin"
	a "github.com/legocy-co/legocy/internal/app"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/admin/marketItem"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/image"
	"github.com/legocy-co/legocy/pkg/auth/jwt/middleware"
)

func AddAdminMarketItems(rg *gin.RouterGroup, app *a.App) {

	marketItemsAdmin := rg.Group("/market-items").Use(middleware.IsAdmin())
	{
		handler := h.NewMarketItemAdminHandler(app.GetMarketItemAdminService())
		{
			marketItemsAdmin.GET("/", handler.GetMarketItemsAdmin)
			marketItemsAdmin.GET("/:itemId", handler.GetMarketItemByID)
			marketItemsAdmin.POST("/", handler.CreateMarketItem)
			marketItemsAdmin.PUT("/:itemId", handler.UpdateMarketItemByID)
			marketItemsAdmin.DELETE("/:itemId", handler.DeleteMarketItemById)
		}
	}

	marketItemImagesAdmin := rg.Group("/market-items/images").Use(middleware.IsAdmin())
	{
		handler := image.NewHandler(app.GetMarketItemImageService(), app.GetImageStorageClient())
		{
			marketItemImagesAdmin.POST("/:marketItemID", handler.UploadImage)
			marketItemImagesAdmin.DELETE("/:imageId", handler.Delete)
		}
	}
}

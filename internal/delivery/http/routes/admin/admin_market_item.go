package admin

import (
	"github.com/gin-gonic/gin"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/admin/marketItem"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/marketplace/image"
	"github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	a "github.com/legocy-co/legocy/internal/pkg/app"
)

func AddAdminMarketItems(rg *gin.RouterGroup, app *a.App) {

	marketItemsAdmin := rg.Group("/market-items").Use(auth.IsAdmin())
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

	marketItemImagesAdmin := rg.Group("/market-items/images").Use(auth.IsAdmin())
	{
		handler := image.NewHandler(app.GetMarketItemImageService(), app.GetImageStorageClient())
		{
			marketItemImagesAdmin.POST("/:marketItemID", handler.UploadImage)
			marketItemImagesAdmin.DELETE("/:imageId", handler.Delete)
		}
	}
}

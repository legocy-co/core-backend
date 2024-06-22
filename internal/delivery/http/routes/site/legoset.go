package site

import (
	"github.com/gin-gonic/gin"
	ih "github.com/legocy-co/legocy/internal/delivery/http/handlers/lego/image"
	h "github.com/legocy-co/legocy/internal/delivery/http/handlers/lego/legoset"
	m "github.com/legocy-co/legocy/internal/delivery/http/middleware/auth"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func AddLegoSets(rg *gin.RouterGroup, app *app.App) {
	handler := h.NewLegoSetHandler(app.GetLegoSetService())
	imagesHandler := ih.NewLegoSetImageHandler(
		app.GetLegoSetImageService(),
		app.GetImageStorageClient(),
	)

	sets := rg.Group("/sets").Use(m.IsAuthenticated())
	{
		sets.GET("/", handler.ListSetsPaginated)
		sets.GET("/all", handler.ListSets)
		sets.GET("/:setID", handler.SetDetail)
	}
	setsAdmin := rg.Group("/admin/sets").Use(m.IsAdmin())
	{
		setsAdmin.POST("/", handler.SetCreate)
		setsAdmin.DELETE("/:setID", handler.SetDelete)
		setsAdmin.PUT("/:setID", handler.SetUpdate)
	}

	// Images
	setsImagesAdmin := rg.Group("/admin/sets/images").Use(m.IsAdmin())
	{
		setsImagesAdmin.DELETE("/:imageId", imagesHandler.DeleteImageById)
		setsImagesAdmin.POST("/:legoSetID", imagesHandler.Upload)
	}
}

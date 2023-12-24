package site

import (
	"github.com/gin-gonic/gin"
	"github.com/legocy-co/legocy/internal/app"
	"github.com/legocy-co/legocy/internal/delivery/http/handlers/utils/image"
)

func AddImagesRoutes(r *gin.RouterGroup, app *app.App) {

	handler := image.NewHandler(app.GetImageStorageClient())

	imgRouter := r.Group("/images")
	{
		imgRouter.GET("/download", handler.DownloadImage)
	}

}

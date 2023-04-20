package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/app"
	"legocy-go/internal/delievery/http/handlers/users/user_image"
	m "legocy-go/internal/delievery/http/middleware"
)

func (r V1router) addUserImages(rg *gin.RouterGroup, app *app.App) {
	handler := user_image.NewUserImageHandler(
		app.GetUserImagesService(),
		app.GetImageStorageClient())

	userImages := rg.Group("/users/images")
	{
		userImages.GET("/:userID", m.UserIdOrAdmin("userID"), handler.ListImages)
		userImages.POST("/:userID", m.UserIdOrAdmin("userID"), handler.UploadUserImage)
		userImages.GET("/download", handler.DownloadImage)
	}
}

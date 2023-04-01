package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/app"
	"legocy-go/internal/delievery/http/handlers/users"
	m "legocy-go/internal/delievery/http/middleware"
)

func (r V1router) addUserImages(rg *gin.RouterGroup, app *app.App) {
	handler := users.NewUserImageHandler(
		app.GetUserImagesService(),
		app.GetImageStorageClient())

	userImages := rg.Group("/users/images/").Use(m.UserIdOrAdmin("userID"))
	{
		userImages.POST("/:userID", handler.UploadUserImage)
	}
}

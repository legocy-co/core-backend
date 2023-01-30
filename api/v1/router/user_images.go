package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/api/v1/handlers/auth"
	m "legocy-go/api/v1/middleware"
	s "legocy-go/api/v1/usecase/auth"
	"legocy-go/internal/storage"
)

func (r V1router) addUserImages(
	rg *gin.RouterGroup, service s.UserImageUseCase, storage storage.ImageStorage) {
	handler := auth.NewUserImageHandler(service, storage)

	userImages := rg.Group("/users/images/").Use(m.UserIdOrAdmin("userID"))
	{
		userImages.POST("/:userID", handler.UploadUserImage)
	}
}

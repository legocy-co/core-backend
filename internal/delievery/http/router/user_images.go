package v1

import (
	"github.com/gin-gonic/gin"
	"legocy-go/internal/delievery/http/handlers/auth"
	m "legocy-go/internal/delievery/http/middleware"
	s "legocy-go/internal/delievery/http/service/auth"
	"legocy-go/internal/storage"
)

func (r V1router) addUserImages(
	rg *gin.RouterGroup, service s.UserImageUseCase, storage storage.ImageStorage) {
	handler := auth.NewUserImageHandler(service, storage)

	userImages := rg.Group("/users/images/").Use(m.Auth())
	{
		userImages.POST("/:userID", handler.UploadUserImage)
	}
}

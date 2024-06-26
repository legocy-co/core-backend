package userImage

import (
	"github.com/legocy-co/legocy/internal/domain/users/service"
	"github.com/legocy-co/legocy/internal/pkg/s3/client"
)

type UserImageHandler struct {
	service service.UserImageService
	storage client.ImageStorage
}

func NewUserImageHandler(
	service service.UserImageService,
	storage client.ImageStorage) UserImageHandler {

	return UserImageHandler{
		service: service,
		storage: storage,
	}
}

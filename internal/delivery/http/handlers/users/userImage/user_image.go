package userImage

import (
	"github.com/legocy-co/legocy/internal/domain/users/service"
	"github.com/legocy-co/legocy/pkg/storage/client"
)

type UserImageHandler struct {
	service service.UserImageUseCase
	storage client.ImageStorage
}

func NewUserImageHandler(
	service service.UserImageUseCase,
	storage client.ImageStorage) UserImageHandler {

	return UserImageHandler{
		service: service,
		storage: storage,
	}
}

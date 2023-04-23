package userImage

import (
	service "legocy-go/internal/delievery/http/service/users"
	"legocy-go/pkg/storage/client"
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

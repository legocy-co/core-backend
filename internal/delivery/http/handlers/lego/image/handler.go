package image

import (
	"github.com/legocy-co/legocy/internal/domain/lego/service"
	storage "github.com/legocy-co/legocy/pkg/storage/client"
)

type LegoSetImageHandler struct {
	service service.LegoSetImageService
	storage storage.ImageStorage
}

func NewLegoSetImageHandler(service service.LegoSetImageService, imageStorage storage.ImageStorage) LegoSetImageHandler {
	return LegoSetImageHandler{service: service, storage: imageStorage}
}

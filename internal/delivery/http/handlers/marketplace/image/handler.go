package image

import (
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	"github.com/legocy-co/legocy/internal/pkg/s3/client"
)

type Handler struct {
	service s.MarketItemImageService
	storage client.ImageStorage
}

func NewHandler(
	service s.MarketItemImageService,
	storage client.ImageStorage) Handler {

	return Handler{
		service: service,
		storage: storage,
	}
}

package market_item_image

import (
	s "github.com/legocy-co/legocy/internal/domain/marketplace/service"
	"github.com/legocy-co/legocy/pkg/storage/client"
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

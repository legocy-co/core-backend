package marketItem

import (
	s "legocy-go/internal/domain/marketplace/service/admin"
)

type Handler struct {
	service s.MarketItemAdminService
}

func NewMarketItemAdminHandler(
	service s.MarketItemAdminService) Handler {

	return Handler{
		service: service,
	}
}

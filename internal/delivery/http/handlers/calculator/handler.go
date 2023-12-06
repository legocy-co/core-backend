package calculator

import "github.com/legocy-co/legocy/internal/domain/calculator/service"

type LegoSetValuationHandler struct {
	service service.LegoSetValuationService
}

func NewLegoSetValuationHandler(service service.LegoSetValuationService) LegoSetValuationHandler {
	return LegoSetValuationHandler{
		service: service,
	}
}

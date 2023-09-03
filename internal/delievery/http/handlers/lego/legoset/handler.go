package legoset

import (
	s "legocy-go/internal/domain/lego/service"
)

type LegoSetHandler struct {
	service s.LegoSetUseCase
}

func NewLegoSetHandler(service s.LegoSetUseCase) LegoSetHandler {
	return LegoSetHandler{service: service}
}

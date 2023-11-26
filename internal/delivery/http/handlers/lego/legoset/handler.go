package legoset

import (
	s "legocy-go/internal/domain/lego/service"
)

type LegoSetHandler struct {
	service s.LegoSetService
}

func NewLegoSetHandler(service s.LegoSetService) LegoSetHandler {
	return LegoSetHandler{service: service}
}

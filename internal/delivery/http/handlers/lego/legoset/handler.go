package legoset

import (
	s "github.com/legocy-co/legocy/internal/domain/lego/service"
)

type LegoSetHandler struct {
	service s.LegoSetService
}

func NewLegoSetHandler(service s.LegoSetService) LegoSetHandler {
	return LegoSetHandler{service: service}
}

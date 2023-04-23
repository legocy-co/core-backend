package legoset

import s "legocy-go/internal/delievery/http/service/lego"

type LegoSetHandler struct {
	service s.LegoSetUseCase
}

func NewLegoSetHandler(service s.LegoSetUseCase) LegoSetHandler {
	return LegoSetHandler{service: service}
}

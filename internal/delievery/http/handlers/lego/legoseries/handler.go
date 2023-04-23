package legoseries

import (
	service "legocy-go/internal/delievery/http/service/lego"
)

type LegoSeriesHandler struct {
	service service.LegoSeriesService
}

func NewLegoSeriesHandler(service service.LegoSeriesService) LegoSeriesHandler {
	return LegoSeriesHandler{service: service}
}

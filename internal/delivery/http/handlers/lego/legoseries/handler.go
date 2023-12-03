package legoseries

import (
	"github.com/legocy-co/legocy/internal/domain/lego/service"
)

type LegoSeriesHandler struct {
	service service.LegoSeriesService
}

func NewLegoSeriesHandler(service service.LegoSeriesService) LegoSeriesHandler {
	return LegoSeriesHandler{service: service}
}

package lego

import (
	models "legocy-go/pkg/lego/models"
)

type LegoSeriesRequest struct {
	Name string `json:"name"`
}

func (sr *LegoSeriesRequest) ToLegoSeriesBasic() *models.LegoSeriesBasic {
	return &models.LegoSeriesBasic{
		Name: sr.Name,
	}
}

type LegoSeriesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetLegoSeriesResponse(m *models.LegoSeries) LegoSeriesResponse {
	return LegoSeriesResponse{
		ID:   m.ID,
		Name: m.Name,
	}
}

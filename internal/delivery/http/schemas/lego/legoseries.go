package lego

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSeriesRequest struct {
	Name string `json:"name"`
}

func (sr *LegoSeriesRequest) ToLegoSeriesValueObject() *models.LegoSeriesValueObject {
	return &models.LegoSeriesValueObject{
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

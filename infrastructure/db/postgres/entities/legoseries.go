package postgres

import (
	"legocy-go/pkg/lego/models"
)

type LegoSeriesPostgres struct {
	Name string
}

func (s *LegoSeriesPostgres) ToLegoSeries() *models.LegoSeries {
	return &models.LegoSeries{
		ID:   -1,
		Name: s.Name,
	}
}

package postgres

import (
	models "legocy-go/pkg/lego/models"
)

type LegoSetPostgres struct {
	Model
	Number   int    `gorm:"unique"`
	Name     string `gorm:"unique"`
	NPieces  int
	SeriesID int `gorm:"not null"`
}

func FromLegoSet(s *models.LegoSet) *LegoSetPostgres {
	return &LegoSetPostgres{
		Number:   s.Number,
		Name:     s.Name,
		NPieces:  s.NPieces,
		SeriesID: s.Series.ID,
	}
}

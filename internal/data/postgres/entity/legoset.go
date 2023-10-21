package postgres

import (
	models "legocy-go/internal/domain/lego/models"
)

type LegoSetPostgres struct {
	Model
	Number               int    `gorm:"unique"`
	Name                 string `gorm:"unique"`
	NPieces              int
	LegoSeriesPostgresID uint
	LegoSeries           LegoSeriesPostgres `gorm:"ForeignKey:LegoSeriesPostgresID"`
}

func (lsp *LegoSetPostgres) ToLegoSet() *models.LegoSet {
	return &models.LegoSet{
		ID:      int(lsp.ID),
		Number:  lsp.Number,
		Name:    lsp.Name,
		NPieces: lsp.NPieces,
		Series:  *lsp.LegoSeries.ToLegoSeries(),
	}
}

func FromLegoSet(s *models.LegoSet) *LegoSetPostgres {
	return &LegoSetPostgres{
		Number:               s.Number,
		Name:                 s.Name,
		NPieces:              s.NPieces,
		LegoSeriesPostgresID: uint(s.Series.ID),
	}
}

func FromLegoSetValueObject(s *models.LegoSetValueObject) *LegoSetPostgres {
	return &LegoSetPostgres{
		Number:               s.Number,
		Name:                 s.Name,
		NPieces:              s.NPieces,
		LegoSeriesPostgresID: uint(s.SeriesID),
	}
}

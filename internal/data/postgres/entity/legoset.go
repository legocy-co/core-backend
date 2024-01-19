package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetPostgres struct {
	Model
	Number       int    `gorm:"unique"`
	Name         string `gorm:"unique"`
	NPieces      int
	LegoSeriesID uint                    `gorm:"index:idx_lego_set_lego_series"`
	LegoSeries   LegoSeriesPostgres      `gorm:"ForeignKey:LegoSeriesID"`
	Images       []*LegoSetImagePostgres `gorm:"foreignKey:LegoSetID;constraint:OnDelete:CASCADE;"`
}

func (lsp LegoSetPostgres) TableName() string {
	return "lego_sets"
}

func (lsp *LegoSetPostgres) ToLegoSet() *models.LegoSet {

	images := make([]*models.LegoSetImage, 0, len(lsp.Images))
	for _, image := range lsp.Images {
		images = append(images, image.ToLegoSetImage())
	}

	return &models.LegoSet{
		ID:      int(lsp.ID),
		Number:  lsp.Number,
		Name:    lsp.Name,
		NPieces: lsp.NPieces,
		Series:  *lsp.LegoSeries.ToLegoSeries(),
		Images:  images,
	}
}

func FromLegoSet(s *models.LegoSet) *LegoSetPostgres {
	return &LegoSetPostgres{
		Number:       s.Number,
		Name:         s.Name,
		NPieces:      s.NPieces,
		LegoSeriesID: uint(s.Series.ID),
	}
}

func FromLegoSetValueObject(s *models.LegoSetValueObject) *LegoSetPostgres {
	return &LegoSetPostgres{
		Number:       s.Number,
		Name:         s.Name,
		NPieces:      s.NPieces,
		LegoSeriesID: uint(s.SeriesID),
	}
}

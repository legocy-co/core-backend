package filters

import (
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
	"gorm.io/gorm"
)

func AddLegoSetFilters(db *gorm.DB, criteria domain.LegoSetFilterCriteria) *gorm.DB {

	if criteria.NpiecesGTE != nil {
		db = db.Where("n_pieces >= ?", *criteria.NpiecesGTE)
	}

	if criteria.NpiecesLTE != nil {
		db = db.Where("n_pieces <= ?", *criteria.NpiecesLTE)
	}

	if criteria.SeriesIDs != nil {
		db = db.Where("lego_series_id IN ?", *criteria.SeriesIDs)
	}

	if criteria.SetNumbers != nil {
		db = db.Where("number IN ?", *criteria.SetNumbers)
	}

	if criteria.Name != nil {
		db = db.Where("name ILIKE ?", *criteria.Name)
	}

	return db

}

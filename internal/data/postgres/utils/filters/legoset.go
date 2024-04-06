package filters

import (
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
	"gorm.io/gorm"
	"strings"
)

func AddLegoSetFilters(
	db *gorm.DB,
	criteria *domain.LegoSetFilterCriteria,
	isNested bool,
) *gorm.DB {

	if criteria == nil {
		return db
	}

	if criteria.NpiecesGTE != nil {
		if !isNested {
			db = db.Where("n_pieces >= ?", *criteria.NpiecesGTE)
		} else {
			db = db.Where("lego_sets.n_pieces >= ?", *criteria.NpiecesGTE)
		}
	}

	if criteria.NpiecesLTE != nil {
		if !isNested {
			db = db.Where("n_pieces <= ?", *criteria.NpiecesLTE)
		} else {
			db = db.Where("lego_sets.n_pieces <= ?", *criteria.NpiecesLTE)
		}
	}

	if criteria.SeriesIDs != nil {
		if !isNested {
			db = db.Where("lego_series_id IN ?", criteria.SeriesIDs)
		} else {
			db = db.Where("lego_sets.lego_series_id IN ?", criteria.SeriesIDs)
		}
	}

	if criteria.SetNumbers != nil {
		if !isNested {
			db = db.Where("number IN ?", criteria.SetNumbers)
		} else {
			db = db.Where("lego_sets.number IN ?", criteria.SetNumbers)
		}
	}

	if criteria.Name != nil {
		if !isNested {
			db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*criteria.Name)+"%")
		} else {
			db = db.Where("LOWER(lego_sets.name) LIKE ?", "%"+strings.ToLower(*criteria.Name)+"%")
		}
	}

	return db

}

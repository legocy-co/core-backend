package filters

import (
	"fmt"
	domain "github.com/legocy-co/legocy/internal/domain/lego/filters"
	"gorm.io/gorm"
	"strings"
)

func AddLegoSetFilters(
	db *gorm.DB,
	criteria *domain.LegoSetFilterCriteria,
	isNested bool,
	tablePrefix string,
) *gorm.DB {

	if criteria == nil {
		return db
	}

	if criteria.NpiecesGTE != nil {
		if !isNested {
			db = db.Where("n_pieces >= ?", *criteria.NpiecesGTE)
		} else {
			db = db.Where(fmt.Sprintf("%sn_pieces >= ?", tablePrefix), *criteria.NpiecesGTE)
		}
	}

	if criteria.NpiecesLTE != nil {
		if !isNested {
			db = db.Where("n_pieces <= ?", *criteria.NpiecesLTE)
		} else {
			db = db.Where(fmt.Sprintf("%sn_pieces <= ?", tablePrefix), *criteria.NpiecesLTE)
		}
	}

	if criteria.SeriesIDs != nil {
		if !isNested {
			db = db.Where("lego_series_id IN ?", criteria.SeriesIDs)
		} else {
			db = db.Where(fmt.Sprintf("%slego_series_id IN ?", tablePrefix), criteria.SeriesIDs)
		}
	}

	if criteria.SetNumbers != nil {
		if !isNested {
			db = db.Where("number IN ?", criteria.SetNumbers)
		} else {
			db = db.Where(fmt.Sprintf("%snumber IN ?", tablePrefix), criteria.SetNumbers)
		}
	}

	if criteria.Name != nil {
		if !isNested {
			db = db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(*criteria.Name)+"%")
		} else {
			db = db.Where(fmt.Sprintf("LOWER(%sname) LIKE ?", tablePrefix), "%"+strings.ToLower(*criteria.Name)+"%")
		}
	}

	if criteria.ReleaseYears != nil && len(criteria.ReleaseYears) > 0 {
		if !isNested {
			db = db.Where("release_year IN ?", criteria.ReleaseYears)
		} else {
			db = db.Where(fmt.Sprintf("%srelease_year IN ?", tablePrefix), criteria.ReleaseYears)
		}
	}

	return db

}

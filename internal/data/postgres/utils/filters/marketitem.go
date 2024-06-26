package filters

import (
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	"gorm.io/gorm"
)

func AddMarketItemsFilters(
	db *gorm.DB,
	criteria *domain.MarketItemFilterCriteria,
	isNested bool,
) *gorm.DB {

	if criteria == nil {
		return db
	}

	if criteria.Ids != nil && len(criteria.Ids) > 0 {
		if isNested {
			db = db.Where("market_items.id IN ?", criteria.Ids)
		} else {
			db = db.Where("market_items.id IN ?", criteria.Ids)
		}
	}

	if criteria.SetIds != nil && len(criteria.SetIds) > 0 {
		if isNested {
			db = db.Where("market_items.lego_set_postgres_id IN ?", criteria.SetIds)
		} else {
			db = db.Where("lego_set_postgres_id IN ?", criteria.SetIds)
		}
	}

	if criteria.MinPrice != nil {
		if isNested {
			db = db.Where("market_items.price >= ?", *criteria.MinPrice)
		} else {
			db = db.Where("price >= ?", *criteria.MinPrice)
		}
	}

	if criteria.MaxPrice != nil {
		if isNested {
			db = db.Where("market_items.price <= ?", *criteria.MaxPrice)
		} else {
			db = db.Where("price <= ?", *criteria.MaxPrice)
		}
	}

	if criteria.SetStates != nil && len(criteria.SetStates) > 0 {
		if isNested {
			db = db.Where("market_items.set_state IN ?", criteria.SetStates)
		} else {
			db = db.Where("set_state IN ?", criteria.SetStates)
		}
	}

	if criteria.Locations != nil && len(criteria.Locations) > 0 {
		if isNested {
			db = db.Where("market_items.location IN ?", criteria.Locations)
		} else {
			db = db.Where("location IN ?", criteria.Locations)
		}
	}

	return db
}

package filters

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	legoFilters "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego/filters"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
)

type MarketItemFilterDTO struct {
	PriceGTE  *float64                      `form:"price_gte" json:"price_gte" binding:"queryparam"`
	PriceLTE  *float64                      `form:"price_lte" json:"price_lte" binding:"queryparam"`
	SetStates []string                      `form:"set_state__in" json:"set_state__in" binding:"queryparam"`
	Locations []string                      `form:"location__in" json:"location__in" binding:"queryparam"`
	LegoSet   *legoFilters.LegoSetFilterDTO `form:"lego_set" json:"lego_set" binding:"queryparam"`
}

func (dto *MarketItemFilterDTO) ToCriteria() (*domain.MarketItemFilterCriteria, *errors.AppError) {

	if dto.LegoSet == nil {
		return domain.NewMarketItemFilterCriteria(
			dto.PriceGTE,
			dto.PriceLTE,
			dto.SetStates,
			dto.Locations,
			nil,
		)
	}

	return domain.NewMarketItemFilterCriteria(
		dto.PriceGTE,
		dto.PriceLTE,
		dto.SetStates,
		dto.Locations,
		dto.LegoSet.ToCriteria(),
	)
}

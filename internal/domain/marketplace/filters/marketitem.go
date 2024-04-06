package filters

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	legoErrors "github.com/legocy-co/legocy/internal/domain/lego"
	legoFilters "github.com/legocy-co/legocy/internal/domain/lego/filters"
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type MarketItemFilterCriteria struct {
	SetIds    []int
	MinPrice  *float64
	MaxPrice  *float64
	SetStates []string
	Locations []string
	LegoSet   *legoFilters.LegoSetFilterCriteria
}

func NewMarketItemFilterCriteria(
	setIDs []int,
	minPrice *float64,
	maxPrice *float64,
	setStates []string,
	locations []string,
	legoSet *legoFilters.LegoSetFilterCriteria,
) (*MarketItemFilterCriteria, *errors.AppError) {

	for _, state := range setStates {
		if !lego.IsValidSetState(state) {
			return nil, &legoErrors.ErrInvalidLegoState
		}
	}

	return &MarketItemFilterCriteria{
		SetIds:    setIDs,
		MinPrice:  minPrice,
		MaxPrice:  maxPrice,
		SetStates: setStates,
		Locations: locations,
		LegoSet:   legoSet,
	}, nil
}

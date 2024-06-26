package filters

import (
	legoErrors "github.com/legocy-co/legocy/internal/domain/lego"
	legoFilters "github.com/legocy-co/legocy/internal/domain/lego/filters"
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

type MarketItemFilterCriteria struct {
	Ids       []int
	SetIds    []int
	MinPrice  *float64
	MaxPrice  *float64
	SetStates []string
	Locations []string
	LegoSet   *legoFilters.LegoSetFilterCriteria
}

func NewMarketItemFilterCriteria(
	ids []int,
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

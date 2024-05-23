package marketplace

import (
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type MarketItemImageValueObject struct {
	MarketItemID int    `validate:"required"`
	ImageURL     string `validate:"required"`
	SortIndex    int    `validate:"default=0"`
}

type MarketItemImagePartialVO struct {
	SortIndex int `validate:"required"`
}

func NewMarketItemImageValueObject(marketItemID int, imageURL string, sortIndex int) (*MarketItemImageValueObject, *errors.AppError) {
	return &MarketItemImageValueObject{
		MarketItemID: marketItemID,
		ImageURL:     imageURL,
		SortIndex:    sortIndex,
	}, nil
}

type MarketItemImage struct {
	ID           int
	MarketItemID int    `validate:"required"`
	ImageURL     string `validate:"required"`
	SortIndex    int    `validate:"default=0"`
}

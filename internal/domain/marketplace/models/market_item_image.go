package marketplace

import (
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type MarketItemImageValueObject struct {
	MarketItemID int    `validate:"required"`
	ImageURL     string `validate:"required"`
	IsMain       bool   `validate:"default=false"`
}

func NewMarketItemImageValueObject(marketItemID int, imageURL string, isMain bool) (*MarketItemImageValueObject, *errors.AppError) {
	return &MarketItemImageValueObject{
		MarketItemID: marketItemID,
		ImageURL:     imageURL,
		IsMain:       isMain,
	}, nil
}

type MarketItemImage struct {
	ID           int
	MarketItemID int    `validate:"required"`
	ImageURL     string `validate:"required"`
	IsMain       bool   `validate:"default=false"`
}

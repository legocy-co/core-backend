package marketplace

import "github.com/legocy-co/legocy/internal/app/errors"

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

func NewMarketItemImage(id int, marketItemID int, imageURL string, isMain bool) (*MarketItemImage, *errors.AppError) {
	return &MarketItemImage{
		ID:           id,
		MarketItemID: marketItemID,
		ImageURL:     imageURL,
		IsMain:       isMain,
	}, nil
}

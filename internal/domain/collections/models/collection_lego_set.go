package models

import (
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

var (
	ErrInvalidPurchasePrice = errors.NewAppError(
		errors.ValidationError,
		"Buy price must be greater than 0",
	)
)

type CollectionLegoSet struct {
	ID           int
	LegoSet      lego.LegoSet
	CurrentState string // lego.SetStateBrandNew, etc.
	BuyPrice     float32
}

func NewCollectionLegoSet(legoSet lego.LegoSet, currentState string, buyPrice float32) (*CollectionLegoSet, *errors.AppError) {

	if (buyPrice) <= 0 {
		return nil, &ErrInvalidPurchasePrice
	}

	return &CollectionLegoSet{
		LegoSet:      legoSet,
		CurrentState: currentState,
		BuyPrice:     buyPrice,
	}, nil
}

type CollectionLegoSetValueObject struct {
	LegoSetID    int
	CurrentState string
	BuyPrice     float32
}

package models

import (
	lego "legocy-go/internal/domain/lego/models"
	marketplace "legocy-go/internal/domain/marketplace/models"
)

type CollectionLegoSet struct {
	ID           int
	LegoSet      lego.LegoSet
	CurrentState string // lego.SetStateBrandNew, etc.
	BuyPrice     float32
	Currency     marketplace.Currency
}

type CollectionLegoSetValueObject struct {
	LegoSetID    int
	CurrentState string
	BuyPrice     float32
	CurrencyID   int
}

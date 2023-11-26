package models

import (
	lego "legocy-go/internal/domain/lego/models"
)

type CollectionLegoSet struct {
	ID           int
	LegoSet      lego.LegoSet
	CurrentState string // lego.SetStateBrandNew, etc.
	BuyPrice     float32
}

type CollectionLegoSetValueObject struct {
	LegoSetID    int
	CurrentState string
	BuyPrice     float32
}

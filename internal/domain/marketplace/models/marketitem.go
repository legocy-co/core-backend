package marketplace

import (
	"legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/users/models"
)

type MarketItem struct {
	ID          int
	LegoSet     lego.LegoSet
	Seller      auth.User
	Price       float32
	Location    string
	SetState    string // lego.SetStateBrandNew / etc.
	Status      string // lego.ListingStatusCheckRequired / lego.ListingStatusActive / etc.
	Description string
}

type MarketItemValueObject struct {
	LegoSetID   int
	SellerID    int
	Price       float32
	CurrencyID  int
	Location    string
	Status      string
	SetState    string // SetStateBrandNew / etc.
	Description string
}

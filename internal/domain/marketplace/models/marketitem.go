package marketplace

import (
	"legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/users/models"
)

type MarketItem struct {
	ID       int
	LegoSet  lego.LegoSet
	Seller   auth.User
	Price    float32
	Currency Currency
	Location Location
	SetState string // SetStateBrandNew / etc.
	Status   string // ListingStatusCheckRequired / ListingStatusActive / etc.
}

type MarketItemValueObject struct {
	LegoSetID  int
	SellerID   int
	Price      float32
	CurrencyID int
	LocationID int
	Status     string
	SetState   string // SetStateBrandNew / etc.
}

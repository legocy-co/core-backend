package marketplace

import (
	lego "legocy-go/internal/domain/lego/models"
	auth "legocy-go/internal/domain/users/models"
)

type MarketItemAdminValueObject struct {
	LegoSetID  int
	SellerID   int
	Price      float32
	CurrencyID int
	LocationID int
}

type MarketItemAdmin struct {
	ID       int
	LegoSet  lego.LegoSet
	Seller   auth.User
	Price    float32
	Currency Currency
	Location Location
}

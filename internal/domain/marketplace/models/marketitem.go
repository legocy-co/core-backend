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
}

type MarketItemBasic struct {
	LegoSetID  int
	SellerID   int
	Price      float32
	CurrencyID int
	LocationID int
}

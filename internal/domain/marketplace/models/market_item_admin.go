package marketplace

import (
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	auth "github.com/legocy-co/legocy/internal/domain/users/models"
)

type MarketItemAdminValueObject struct {
	LegoSetID   int
	SellerID    int
	Price       float32
	Location    string
	Status      string
	SetState    string
	Description string
}

type MarketItemAdmin struct {
	ID          int
	LegoSet     lego.LegoSet
	Seller      auth.User
	Price       float32
	Location    string
	Status      string
	SetState    string
	Description string
	Images      []*MarketItemImage
}

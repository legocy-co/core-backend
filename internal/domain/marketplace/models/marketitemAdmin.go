package marketplace

import (
	lego "legocy-go/internal/domain/lego/models"
	auth "legocy-go/internal/domain/users/models"
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
}

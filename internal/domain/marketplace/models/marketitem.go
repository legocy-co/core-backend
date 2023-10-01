package marketplace

import (
	"legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/users/models"
)

const (
	CheckRequired = "CHECK_REQUIRED"
	Active        = "ACTIVE"
)

func IsValidStatus(status string) bool {
	return status == CheckRequired || status == Active
}

type MarketItem struct {
	ID       int
	LegoSet  lego.LegoSet
	Seller   auth.User
	Price    float32
	Currency Currency
	Location Location
	Status   string // CheckRequired / Active / etc.
}

type MarketItemValueObject struct {
	LegoSetID  int
	SellerID   int
	Price      float32
	CurrencyID int
	LocationID int
	Status     string
}

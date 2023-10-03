package marketplace

import (
	"legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/users/models"
)

const (
	ListingStatusCheckRequired = "CHECK_REQUIRED"
	ListingStatusActive        = "ACTIVE"
	ListingStatusSold          = "SOLD"

	SetStateBrandNew        = "BRAND_NEW"
	SetStateBoxOpened       = "BOX_OPENED"
	SetStateBagsOpened      = "BAGS_OPENED"
	SetStateBuiltWithBox    = "BUILT_WITH_BOX"
	SetStateBuiltWithoutBox = "BUILT_WITHOUT_BOX"
)

func IsValidListingStatus(status string) bool {
	validStatuses := [3]string{
		ListingStatusCheckRequired,
		ListingStatusActive,
		ListingStatusSold,
	}

	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}

	return false

}

type MarketItem struct {
	ID       int
	LegoSet  lego.LegoSet
	Seller   auth.User
	Price    float32
	Currency Currency
	Location Location
	Status   string // ListingStatusCheckRequired / ListingStatusActive / etc.
}

type MarketItemValueObject struct {
	LegoSetID  int
	SellerID   int
	Price      float32
	CurrencyID int
	LocationID int
	Status     string
}

package marketplace

import (
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	users "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

var (
	ErrInvalidListingStatus = errors.NewAppError(errors.ValidationError, "invalid listing status")
	ErrInvalidSetState      = errors.NewAppError(errors.ValidationError, "invalid set state")
)

type MarketItem struct {
	ID          int
	LegoSet     lego.LegoSet
	Seller      users.User
	Price       float32
	Location    string
	SetState    string // lego.SetStateBrandNew / etc.
	Status      string // lego.ListingStatusCheckRequired / lego.ListingStatusActive / etc.
	Description string
	Images      []*MarketItemImage
	Liked       bool
}

func NewMarketItem(
	id int,
	legoSet lego.LegoSet,
	seller users.User,
	price float32,
	location string,
	setState string,
	status string,
	description string,
	images []*MarketItemImage,
	liked bool,
) (*MarketItem, *errors.AppError) {

	if !IsValidListingStatus(status) {
		return nil, &ErrInvalidListingStatus
	}

	if !lego.IsValidSetState(setState) {
		return nil, &ErrInvalidSetState
	}

	return &MarketItem{
		ID:          id,
		LegoSet:     legoSet,
		Seller:      seller,
		Price:       price,
		Location:    location,
		SetState:    setState,
		Status:      status,
		Description: description,
		Images:      images,
		Liked:       liked,
	}, nil
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

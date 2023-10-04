package admin

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminCreateRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	SellerID    int     `json:"seller_id"`
	Price       float32 `json:"price"`
	CurrencyID  int     `json:"currency_id"`
	LocationID  int     `json:"location_id"`
	Status      string  `json:"status"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r MarketItemAdminCreateRequest) ToMarketItemAdminValueObject() (
	*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	if !models.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    r.SellerID,
		Price:       r.Price,
		CurrencyID:  r.CurrencyID,
		LocationID:  r.LocationID,
		Status:      r.Status,
		SetState:    r.SetState,
		Description: r.Description,
	}, nil
}

type MarketItemAdminUpdateRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	SellerID    int     `json:"seller_id"`
	Price       float32 `json:"price"`
	CurrencyID  int     `json:"currency_id"`
	LocationID  int     `json:"location_id"`
	Status      string  `json:"status"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r MarketItemAdminUpdateRequest) ToMarketItemAdminValueObject() (
	*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	if !models.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    r.SellerID,
		Price:       r.Price,
		CurrencyID:  r.CurrencyID,
		LocationID:  r.LocationID,
		Status:      r.Status,
		SetState:    r.SetState,
		Description: r.Description,
	}, nil
}

type MarketItemAdminResponse struct {
	ID          int                          `json:"id"`
	Price       float32                      `json:"price"`
	Currency    marketplace.CurrencyResponse `json:"currency"`
	Location    marketplace.LocationResponse `json:"location"`
	LegoSet     lego.LegoSetResponse         `json:"lego_set"`
	Seller      users.UserDetailResponse     `json:"seller"`
	Status      string                       `json:"status"`
	SetState    string                       `json:"set_state"`
	Description string                       `json:"description"`
}

func GetMarketItemAdminResponse(
	mi *models.MarketItemAdmin) MarketItemAdminResponse {
	return MarketItemAdminResponse{
		ID:          mi.ID,
		Price:       mi.Price,
		Currency:    marketplace.GetCurrencyResponse(&mi.Currency),
		Location:    marketplace.GetLocationResponse(&mi.Location),
		LegoSet:     lego.GetLegoSetResponse(&mi.LegoSet),
		Seller:      users.GetUserDetailResponse(&mi.Seller),
		Status:      mi.Status,
		SetState:    mi.SetState,
		Description: mi.Description,
	}
}

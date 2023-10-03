package admin

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminCreateRequest struct {
	models.MarketItemAdminValueObject
}

func (r MarketItemAdminCreateRequest) ToMarketItemAdminValueObject() (*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:  r.LegoSetID,
		SellerID:   r.SellerID,
		Price:      r.Price,
		CurrencyID: r.CurrencyID,
		LocationID: r.LocationID,
		Status:     r.Status,
	}, nil
}

type MarketItemAdminUpdateRequest struct {
	models.MarketItemAdminValueObject
}

func (r MarketItemAdminUpdateRequest) ToMarketItemAdminValueObject() (*models.MarketItemAdminValueObject, error) {

	if !models.IsValidListingStatus(r.Status) {
		return nil, errors.ErrMarketItemInvalidStatus
	}

	return &models.MarketItemAdminValueObject{
		LegoSetID:  r.LegoSetID,
		SellerID:   r.SellerID,
		Price:      r.Price,
		CurrencyID: r.CurrencyID,
		LocationID: r.LocationID,
		Status:     r.Status,
	}, nil
}

type MarketItemAdminResponse struct {
	ID       int                          `json:"id"`
	Price    float32                      `json:"price"`
	Currency marketplace.CurrencyResponse `json:"currency"`
	Location marketplace.LocationResponse `json:"location"`
	LegoSet  lego.LegoSetResponse         `json:"lego_set"`
	Seller   users.UserDetailResponse     `json:"seller"`
	Status   string                       `json:"status"`
}

func GetMarketItemAdminResponse(mi *models.MarketItemAdmin) MarketItemAdminResponse {
	return MarketItemAdminResponse{
		ID:       mi.ID,
		Price:    mi.Price,
		Currency: marketplace.GetCurrencyResponse(&mi.Currency),
		Location: marketplace.GetLocationResponse(&mi.Location),
		LegoSet:  lego.GetLegoSetResponse(&mi.LegoSet),
		Seller:   users.GetUserDetailResponse(&mi.Seller),
		Status:   mi.Status,
	}
}

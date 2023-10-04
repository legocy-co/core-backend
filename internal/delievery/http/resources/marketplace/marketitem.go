package marketplace

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/users"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemRequest struct {
	LegoSetID   int     `json:"lego_set_id"`
	Price       float32 `json:"price"`
	CurrencyID  int     `json:"currency_id"`
	LocationID  int     `json:"location_id"`
	SetState    string  `json:"set_state"`
	Description string  `json:"description"`
}

func (r *MarketItemRequest) ToMarketItemValueObject(sellerID int) (*models.MarketItemValueObject, error) {

	if !models.IsValidSetState(r.SetState) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.MarketItemValueObject{
		LegoSetID:   r.LegoSetID,
		SellerID:    sellerID,
		Price:       r.Price,
		CurrencyID:  r.CurrencyID,
		LocationID:  r.LocationID,
		Status:      models.ListingStatusCheckRequired,
		Description: r.Description,
	}, nil
}

type MarketItemResponse struct {
	ID          int                      `json:"id"`
	Price       float32                  `json:"price"`
	Currency    CurrencyResponse         `json:"currency"`
	Location    LocationResponse         `json:"location"`
	LegoSet     lego.LegoSetResponse     `json:"lego_set"`
	Seller      users.UserDetailResponse `json:"seller"`
	Status      string                   `json:"status"`
	SetState    string                   `json:"set_state"`
	Description string                   `json:"description"`
}

func GetMarketItemResponse(m *models.MarketItem) MarketItemResponse {
	return MarketItemResponse{
		ID:          m.ID,
		Price:       m.Price,
		Currency:    GetCurrencyResponse(&m.Currency),
		Location:    GetLocationResponse(&m.Location),
		LegoSet:     lego.GetLegoSetResponse(&m.LegoSet),
		Seller:      users.GetUserDetailResponse(&m.Seller),
		Status:      m.Status,
		SetState:    m.SetState,
		Description: m.Description,
	}
}

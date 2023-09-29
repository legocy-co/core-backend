package admin

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminCreateRequest struct {
	models.MarketItemAdminValueObject
}

type MarketItemAdminUpdateRequest struct {
	models.MarketItemAdminValueObject
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

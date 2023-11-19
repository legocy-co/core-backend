package collections

import (
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
	"legocy-go/internal/delievery/http/resources/valuation"
	"legocy-go/internal/domain/calculator/models"
	auth "legocy-go/internal/domain/users/models"
)

type UserCollectionValuationResponse struct {
	User       users.UserDetailResponse             `json:"user"`
	Valuations []valuation.LegoSetValuationResponse `json:"valuations"`
	Currency   marketplace.CurrencyResponse         `json:"currency"`
	Total      float32                              `json:"total"`
}

func FromUserCollectionValuation(
	valuations []models.LegoSetValuation, user auth.User) UserCollectionValuationResponse {

	valuationsResponse := make([]valuation.LegoSetValuationResponse, 0, len(valuations))
	var total float32
	var currency *marketplace.CurrencyResponse

	for _, valuationDomain := range valuations {
		total += valuationDomain.CompanyValuation
		valuationsResponse = append(
			valuationsResponse,
			valuation.FromLegoSetValuation(valuationDomain),
		)

		if currency == nil {
			*currency = marketplace.GetCurrencyResponse(&valuationDomain.Currency)
		}

	}

	return UserCollectionValuationResponse{
		User:       users.GetUserDetailResponse(&user),
		Valuations: valuationsResponse,
		Currency:   *currency,
		Total:      total,
	}

}

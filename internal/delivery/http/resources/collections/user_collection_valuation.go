package collections

import (
	"legocy-go/internal/delivery/http/resources/users"
	"legocy-go/internal/delivery/http/resources/valuation"
	"legocy-go/internal/domain/calculator/models"
	auth "legocy-go/internal/domain/users/models"
)

type UserCollectionValuationResponse struct {
	User       users.UserDetailResponse             `json:"user"`
	Valuations []valuation.LegoSetValuationResponse `json:"valuations"`
	Total      float32                              `json:"total"`
}

func FromUserCollectionValuation(
	valuations []models.LegoSetValuation, user auth.User) UserCollectionValuationResponse {

	valuationsResponse := make([]valuation.LegoSetValuationResponse, 0, len(valuations))
	var total float32

	for _, valuationDomain := range valuations {
		total += valuationDomain.CompanyValuation
		valuationsResponse = append(
			valuationsResponse,
			valuation.FromLegoSetValuation(valuationDomain),
		)
	}

	return UserCollectionValuationResponse{
		User:       users.GetUserDetailResponse(&user),
		Valuations: valuationsResponse,
		Total:      total,
	}

}

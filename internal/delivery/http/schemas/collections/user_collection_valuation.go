package collections

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	auth "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserCollectionValuationResponse struct {
	User       users.UserDetailResponse              `json:"user"`
	Valuations []calculator.LegoSetValuationResponse `json:"valuations"`
	Total      float32                               `json:"total"`
}

func FromUserCollectionValuation(
	valuations []models.LegoSetValuation, user auth.User) UserCollectionValuationResponse {

	valuationsResponse := make([]calculator.LegoSetValuationResponse, 0, len(valuations))
	var total float32

	for _, valuationDomain := range valuations {
		total += valuationDomain.CompanyValuation
		valuationsResponse = append(
			valuationsResponse,
			calculator.FromLegoSetValuation(valuationDomain),
		)
	}

	return UserCollectionValuationResponse{
		User:       users.GetUserDetailResponse(&user),
		Valuations: valuationsResponse,
		Total:      total,
	}

}

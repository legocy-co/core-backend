package collections

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	"github.com/legocy-co/legocy/internal/domain/collections/service/collection/pl"
	auth "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserCollectionValuationResponse struct {
	User       users.UserDetailResponse              `json:"user"`
	Valuations []calculator.LegoSetValuationResponse `json:"valuations"`
	Total      float32                               `json:"total"`
}

type UserCollectionValuationTotals struct {
	Total        float32                        `json:"total"`
	TotalSets    int                            `json:"total_sets"`
	SetsValuated int                            `json:"sets_valuated"`
	TotalProfits CollectionTotalProfitsResponse `json:"total_profits"`
}

type CollectionTotalProfitsResponse struct {
	TotalReturnUSD        float32 `json:"total_return_usd"`
	TotalReturnPercentage float32 `json:"total_return_percentage"`
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

func GetCollectionValuationTotals(collectionSets []pl.SetWithValuation) UserCollectionValuationTotals {

	var total float32
	var setsValuated int

	for _, s := range collectionSets {
		if s.SetValuation != nil {
			total += s.SetValuation.CompanyValuation
			setsValuated++
		}
	}

	return UserCollectionValuationTotals{
		Total:        total,
		TotalSets:    len(collectionSets),
		SetsValuated: setsValuated,
		TotalProfits: GetTotalProfitsResponse(collectionSets),
	}
}

func GetTotalProfitsResponse(collectionSets []pl.SetWithValuation) CollectionTotalProfitsResponse {
	profits := pl.GetCollectionProfits(collectionSets)
	return CollectionTotalProfitsResponse{
		TotalReturnUSD:        profits.TotalReturnUSD,
		TotalReturnPercentage: profits.TotalReturnPercentage,
	}
}

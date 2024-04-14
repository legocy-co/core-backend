package collections

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	collections "github.com/legocy-co/legocy/internal/domain/collections/models"
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
	TotalSets    int                            `json:"totalSets"`
	SetsValuated int                            `json:"setsValuated"`
	TotalProfits CollectionTotalProfitsResponse `json:"totalProfits"`
}

type CollectionTotalProfitsResponse struct {
	TotalReturnUSD        float32 `json:"totalReturnUSD"`
	TotalReturnPercentage float32 `json:"totalReturnPercentage"`
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

func GetCollectionValuationTotals(collectionSets []collections.SetWithValuation) UserCollectionValuationTotals {

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

func GetTotalProfitsResponse(collectionSets []collections.SetWithValuation) CollectionTotalProfitsResponse {
	profits := pl.GetCollectionProfits(collectionSets)
	return CollectionTotalProfitsResponse{
		TotalReturnUSD:        profits.TotalReturnUSD,
		TotalReturnPercentage: profits.TotalReturnPercentage,
	}
}

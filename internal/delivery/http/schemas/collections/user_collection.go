package collections

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	calculatorModels "github.com/legocy-co/legocy/internal/domain/calculator/models"
	"github.com/legocy-co/legocy/internal/domain/collections/models"
)

type UserLegoSetCollectionResponse struct {
	User   users.UserDetailResponse      `json:"user"`
	Sets   []CollectionLegoSetResponse   `json:"collection_sets"`
	Totals UserCollectionValuationTotals `json:"totals"`
}

func GetUserLegoCollectionResponse(
	collection models.LegoCollection,
	valuations []calculatorModels.LegoSetValuation,
) UserLegoSetCollectionResponse {

	setsResponses := make([]CollectionLegoSetResponse, 0, len(collection.Sets))
	var setsWithValuation []models.SetWithValuation

	for _, set := range collection.Sets {

		var setValuation *calculatorModels.LegoSetValuation = nil
		for _, valuation := range valuations {
			if valuation.LegoSet.ID == set.LegoSet.ID && valuation.State == set.CurrentState {
				setValuation = &valuation
				break
			}
		}

		setsResponses = append(setsResponses, GetCollectionLegoSetResponse(set, setValuation))
		setsWithValuation = append(setsWithValuation, models.NewSetWithValuation(set, setValuation))
	}

	collectionTotals := GetCollectionValuationTotals(setsWithValuation)

	return UserLegoSetCollectionResponse{
		User:   users.GetUserDetailResponse(&collection.User),
		Sets:   setsResponses,
		Totals: collectionTotals,
	}
}

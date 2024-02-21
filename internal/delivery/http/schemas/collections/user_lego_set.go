package collections

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/calculator"
	legoResources "github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	calculatorModels "github.com/legocy-co/legocy/internal/domain/calculator/models"
	"github.com/legocy-co/legocy/internal/domain/collections/models"
	"github.com/legocy-co/legocy/internal/domain/collections/service/collection/pl"
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/domain/marketplace/errors"
)

type CollectionLegoSetResponse struct {
	ID         int                                  `json:"id"`
	LegoSet    legoResources.LegoSetResponse        `json:"lego_set"`
	Valuation  *calculator.LegoSetValuationResponse `json:"valuation"`
	SetProfits *UserLegoSetProfitsResponse          `json:"set_profits"`
	State      string                               `json:"state"`
	BuyPrice   float32                              `json:"buy_price"`
}

type UserLegoSetProfitsResponse struct {
	TotalReturnUSD        float32 `json:"total_return_usd"`
	TotalReturnPercentage float32 `json:"total_return_percentage"`
}

func GetCollectionLegoSetResponse(collectionSet models.CollectionLegoSet, valuation *calculatorModels.LegoSetValuation) CollectionLegoSetResponse {

	var valuationResponse *calculator.LegoSetValuationResponse
	if valuation != nil {
		valuationResponseValue := calculator.FromLegoSetValuation(*valuation)
		valuationResponse = &valuationResponseValue
	}

	return CollectionLegoSetResponse{
		ID:         collectionSet.ID,
		LegoSet:    legoResources.GetLegoSetResponse(&collectionSet.LegoSet),
		Valuation:  valuationResponse,
		SetProfits: GetUserLegoSetProfitsResponse(collectionSet, valuation),
		State:      collectionSet.CurrentState,
		BuyPrice:   collectionSet.BuyPrice,
	}
}

func GetUserLegoSetProfitsResponse(
	collectionSet models.CollectionLegoSet,
	valuation *calculatorModels.LegoSetValuation) *UserLegoSetProfitsResponse {

	profits := pl.GetCollectionSetProfits(
		models.NewSetWithValuation(collectionSet, valuation),
	)

	return &UserLegoSetProfitsResponse{
		TotalReturnUSD:        profits.ReturnUSD,
		TotalReturnPercentage: profits.ReturnPercentage,
	}
}

type CollectionLegoSetAddRequest struct {
	LegoSetID int     `json:"lego_set_id"`
	State     string  `json:"state"`
	BuyPrice  float32 `json:"buy_price"`
}

func (r CollectionLegoSetAddRequest) ToCollectionLegoSetValueObject() (*models.CollectionLegoSetValueObject, error) {
	if !lego.IsValidSetState(r.State) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.CollectionLegoSetValueObject{
		LegoSetID:    r.LegoSetID,
		CurrentState: r.State,
		BuyPrice:     r.BuyPrice,
	}, nil
}

type CollectionLegoSetUpdateRequest struct {
	LegoSetID int     `json:"lego_set_id"`
	State     string  `json:"state"`
	BuyPrice  float32 `json:"buy_price"`
}

func (r CollectionLegoSetUpdateRequest) ToCollectionLegoSetValueObject() (*models.CollectionLegoSetValueObject, error) {
	if !lego.IsValidSetState(r.State) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.CollectionLegoSetValueObject{
		LegoSetID:    r.LegoSetID,
		CurrentState: r.State,
		BuyPrice:     r.BuyPrice,
	}, nil
}

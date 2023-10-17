package collections

import (
	legoResources "legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/domain/collections/models"
	lego "legocy-go/internal/domain/lego/models"
	"legocy-go/internal/domain/marketplace/errors"
)

type CollectionLegoSetResponse struct {
	ID       int                           `json:"id"`
	LegoSet  legoResources.LegoSetResponse `json:"lego_set"`
	State    string                        `json:"state"`
	BuyPrice float32                       `json:"buy_price"`
	Currency marketplace.CurrencyResponse  `json:"currency"`
}

func GetCollectionLegoSetResponse(collectionSet models.CollectionLegoSet) CollectionLegoSetResponse {
	return CollectionLegoSetResponse{
		ID:       collectionSet.ID,
		LegoSet:  legoResources.GetLegoSetResponse(&collectionSet.LegoSet),
		State:    collectionSet.CurrentState,
		BuyPrice: collectionSet.BuyPrice,
		Currency: marketplace.GetCurrencyResponse(&collectionSet.Currency),
	}
}

type CollectionLegoSetAddRequest struct {
	LegoSetID  int     `json:"lego_set_id"`
	State      string  `json:"state"`
	BuyPrice   float32 `json:"buy_price"`
	CurrencyID int     `json:"currency_id"`
}

func (r CollectionLegoSetAddRequest) ToCollectionLegoSetValueObject() (*models.CollectionLegoSetValueObject, error) {
	if !lego.IsValidSetState(r.State) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.CollectionLegoSetValueObject{
		LegoSetID:    r.LegoSetID,
		CurrentState: r.State,
		BuyPrice:     r.BuyPrice,
		CurrencyID:   r.CurrencyID,
	}, nil
}

type CollectionLegoSetUpdateRequest struct {
	LegoSetID  int     `json:"lego_set_id"`
	State      string  `json:"state"`
	BuyPrice   float32 `json:"buy_price"`
	CurrencyID int     `json:"currency_id"`
}

func (r CollectionLegoSetUpdateRequest) ToCollectionLegoSetValueObject() (*models.CollectionLegoSetValueObject, error) {
	if !lego.IsValidSetState(r.State) {
		return nil, errors.ErrMarketItemInvalidSetState
	}

	return &models.CollectionLegoSetValueObject{
		LegoSetID:    r.LegoSetID,
		CurrentState: r.State,
		BuyPrice:     r.BuyPrice,
		CurrencyID:   r.CurrencyID,
	}, nil
}

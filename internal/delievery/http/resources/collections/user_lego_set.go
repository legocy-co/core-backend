package collections

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/domain/collections/models"
)

type CollectionLegoSetResponse struct {
	ID       int                          `json:"id"`
	LegoSet  lego.LegoSetResponse         `json:"lego_set"`
	State    string                       `json:"state"`
	BuyPrice float32                      `json:"buy_price"`
	Currency marketplace.CurrencyResponse `json:"currency"`
}

func GetCollectionLegoSetResponse(collectionSet models.CollectionLegoSet) CollectionLegoSetResponse {
	return CollectionLegoSetResponse{
		ID:       collectionSet.ID,
		LegoSet:  lego.GetLegoSetResponse(&collectionSet.LegoSet),
		State:    collectionSet.CurrentState,
		BuyPrice: collectionSet.BuyPrice,
		Currency: marketplace.GetCurrencyResponse(&collectionSet.Currency),
	}
}

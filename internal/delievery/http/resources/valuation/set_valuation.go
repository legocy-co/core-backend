package valuation

import (
	"legocy-go/internal/delievery/http/resources/lego"
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/domain/calculator/models"
)

type LegoSetValuationResponse struct {
	ID        int                          `json:"id"`
	LegoSet   lego.LegoSetResponse         `json:"lego_set"`
	State     string                       `json:"state"`
	Valuation float32                      `json:"valuation"`
	Currency  marketplace.CurrencyResponse `json:"currency"`
}

func FromLegoSetValuation(v models.LegoSetValuation) LegoSetValuationResponse {
	return LegoSetValuationResponse{
		ID:        v.ID,
		LegoSet:   lego.GetLegoSetResponse(&v.LegoSet),
		State:     v.State,
		Valuation: v.CompanyValuation,
		Currency:  marketplace.GetCurrencyResponse(&v.Currency),
	}
}

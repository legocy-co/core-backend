package valuation

import (
	"legocy-go/internal/delivery/http/resources/lego"
	"legocy-go/internal/domain/calculator/models"
)

type LegoSetValuationResponse struct {
	ID        int                  `json:"id"`
	LegoSet   lego.LegoSetResponse `json:"lego_set"`
	State     string               `json:"state"`
	Valuation float32              `json:"valuation"`
}

func FromLegoSetValuation(v models.LegoSetValuation) LegoSetValuationResponse {
	return LegoSetValuationResponse{
		ID:        v.ID,
		LegoSet:   lego.GetLegoSetResponse(&v.LegoSet),
		State:     v.State,
		Valuation: v.CompanyValuation,
	}
}

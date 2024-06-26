package calculator

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/lego"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
)

type LegoSetValuationResponse struct {
	ID        int                  `json:"id"`
	LegoSet   lego.LegoSetResponse `json:"legoSet"`
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

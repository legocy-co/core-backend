package calculator

import (
	"errors"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
	"github.com/legocy-co/legocy/internal/domain/lego"
	legoModels "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetValuationCreateRequest struct {
	LegoSetID int     `json:"lego_set_id"`
	State     string  `json:"state"`
	Valuation float32 `json:"valuation"`
}

func (r LegoSetValuationCreateRequest) ToLegoSetValuationVO() (
	*models.LegoSetValuationValueObject, error) {

	if !legoModels.IsValidSetState(r.State) {
		return nil, lego.ErrInvalidLegoState
	}

	if r.Valuation <= 0 {
		return nil, errors.New("valuation must be a positive number")
	}

	return &models.LegoSetValuationValueObject{
		LegoSetID:        r.LegoSetID,
		State:            r.State,
		CompanyValuation: r.Valuation,
	}, nil
}

type LegoSetValuationUpdateRequest struct {
	LegoSetID int     `json:"lego_set_id"`
	State     string  `json:"state"`
	Valuation float32 `json:"valuation"`
}

func (r LegoSetValuationUpdateRequest) ToLegoSetValuationVO() (
	*models.LegoSetValuationValueObject, error) {

	if !legoModels.IsValidSetState(r.State) {
		return nil, lego.ErrInvalidLegoState
	}

	if r.Valuation <= 0 {
		return nil, errors.New("valuation must be a positive number")
	}

	return &models.LegoSetValuationValueObject{
		LegoSetID:        r.LegoSetID,
		State:            r.State,
		CompanyValuation: r.Valuation,
	}, nil
}

package repository

import (
	"context"
	"legocy-go/internal/domain/collections/models"
)

type LegoSetValuationRepository interface {
	GetLegoSetValuationsList(c context.Context, legoSetID int) ([]models.LegoSetValuation, error)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, error)
	GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, error)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) error
	DeleteLegoSetValuationByID(c context.Context, id int) error
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) error
}

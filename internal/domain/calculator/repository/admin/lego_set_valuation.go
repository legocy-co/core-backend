package admin

import (
	"context"
	"legocy-go/internal/domain/calculator/models"
	"legocy-go/internal/domain/errors"
)

type LegoSetValuationAdminRepository interface {
	GetLegoSetValuations(c context.Context) ([]*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError
	DeleteLegoSetValuationByID(c context.Context, id int) *errors.AppError
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError
}

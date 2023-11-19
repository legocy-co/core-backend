package admin

import (
	"context"
	"legocy-go/internal/app/errors"
	"legocy-go/internal/domain/calculator/models"
)

type LegoSetValuationAdminRepository interface {
	GetLegoSetValuations(c context.Context) ([]*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError
	DeleteLegoSetValuationByID(c context.Context, id int) *errors.AppError
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError
}

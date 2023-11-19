package repository

import (
	"context"
	"legocy-go/internal/app/errors"
	"legocy-go/internal/domain/calculator/models"
)

type LegoSetValuationRepository interface {
	GetLegoSetValuationsList(c context.Context, legoSetID int) ([]models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError)
}

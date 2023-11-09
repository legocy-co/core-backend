package repository

import (
	"context"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/errors"
)

type LegoSetValuationRepository interface {
	GetLegoSetValuationsList(c context.Context, legoSetID int) ([]models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationBySetStateCurrency(c context.Context, setID int, setState string, currencyID int) (*models.LegoSetValuation, *errors.AppError)
}

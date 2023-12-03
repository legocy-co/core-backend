package admin

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	"github.com/legocy-co/legocy/internal/domain/calculator/models"
)

type LegoSetValuationAdminRepository interface {
	GetLegoSetValuations(c context.Context) ([]*models.LegoSetValuation, *errors.AppError)
	GetLegoSetValuationByID(c context.Context, id int) (*models.LegoSetValuation, *errors.AppError)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError
	DeleteLegoSetValuationByID(c context.Context, id int) *errors.AppError
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError
}

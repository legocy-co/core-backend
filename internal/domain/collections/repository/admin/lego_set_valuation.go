package admin

import (
	"context"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/errors"
)

type LegoSetValuationAdminRepository interface {
	GetLegoSetValuations(c context.Context) ([]*models.LegoSetValuation, *errors.AppError)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) *errors.AppError
	DeleteLegoSetValuationByID(c context.Context, id int) *errors.AppError
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) *errors.AppError
}

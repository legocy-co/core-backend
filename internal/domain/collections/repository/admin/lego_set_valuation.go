package admin

import (
	"context"
	"legocy-go/internal/domain/collections/models"
)

type LegoSetValuationAdminRepository interface {
	GetLegoSetValuations(c context.Context) ([]*models.LegoSetValuation, error)
	AddLegoSetValuation(c context.Context, vo models.LegoSetValuationValueObject) error
	DeleteLegoSetValuationByID(c context.Context, id int) error
	UpdateLegoSetValuationByID(c context.Context, id int, vo models.LegoSetValuationValueObject) error
}

package marketplace

import (
	"context"
	"legocy-go/internal/domain/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type LocationRepository interface {
	GetLocations(c context.Context) ([]*models.Location, *errors.AppError)
	GetCountryLocations(c context.Context, country string) ([]*models.Location, *errors.AppError)
	CreateLocation(c context.Context, location *models.LocationValueObject) *errors.AppError
	DeleteLocation(c context.Context, id int) *errors.AppError
}

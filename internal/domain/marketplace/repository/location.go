package marketplace

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
)

type LocationRepository interface {
	GetLocations(c context.Context) ([]*models.Location, error)
	GetCountryLocations(c context.Context, country string) ([]*models.Location, error)
	CreateLocation(c context.Context, location *models.LocationBasic) error
	DeleteLocation(c context.Context, id int) error
}

package marketplace

import (
	"context"
	models "legocy-go/pkg/marketplace/models"
	r "legocy-go/pkg/marketplace/repository"
)

type LocationUseCase struct {
	repo r.LocationRepository
}

func NewLocationUseCase(repo r.LocationRepository) LocationUseCase {
	return LocationUseCase{repo: repo}
}

func (l *LocationUseCase) ListLocations(c context.Context) ([]*models.Location, error) {
	return l.repo.GetLocations(c)
}

func (l *LocationUseCase) CountryLocations(c context.Context, country string) ([]*models.Location, error) {
	return l.repo.GetCountryLocations(c, country)
}

func (l *LocationUseCase) CreateLocation(c context.Context, loc *models.LocationBasic) error {
	return l.repo.CreateLocation(c, loc)
}

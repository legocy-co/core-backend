package service

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
	r "legocy-go/internal/domain/marketplace/repository"
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

func (l *LocationUseCase) CreateLocation(c context.Context, loc *models.LocationValueObject) error {
	return l.repo.CreateLocation(c, loc)
}

func (l *LocationUseCase) DeleteLocation(c context.Context, id int) error {
	return l.repo.DeleteLocation(c, id)
}

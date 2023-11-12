package service

import (
	"context"
	"legocy-go/internal/app/errors"
	models "legocy-go/internal/domain/marketplace/models"
	r "legocy-go/internal/domain/marketplace/repository"
)

type LocationUseCase struct {
	repo r.LocationRepository
}

func NewLocationUseCase(repo r.LocationRepository) LocationUseCase {
	return LocationUseCase{repo: repo}
}

func (l *LocationUseCase) ListLocations(c context.Context) ([]*models.Location, *errors.AppError) {
	return l.repo.GetLocations(c)
}

func (l *LocationUseCase) CountryLocations(c context.Context, country string) ([]*models.Location, *errors.AppError) {
	return l.repo.GetCountryLocations(c, country)
}

func (l *LocationUseCase) CreateLocation(c context.Context, loc *models.LocationValueObject) *errors.AppError {
	return l.repo.CreateLocation(c, loc)
}

func (l *LocationUseCase) DeleteLocation(c context.Context, id int) *errors.AppError {
	return l.repo.DeleteLocation(c, id)
}

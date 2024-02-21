package repository

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSeriesRepository interface {
	CreateLegoSeries(c context.Context, s *models.LegoSeriesValueObject) *errors.AppError
	GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, *errors.AppError)
	GetLegoSeries(c context.Context, id int) (*models.LegoSeries, *errors.AppError)
	GetLegoSeriesByName(c context.Context, name string) (*models.LegoSeries, *errors.AppError)
	DeleteLegoSeries(c context.Context, id int) *errors.AppError
	UpdateLegoSeries(id int, s *models.LegoSeriesValueObject) *errors.AppError
}

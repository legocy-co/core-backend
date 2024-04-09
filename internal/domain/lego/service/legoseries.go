package service

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	r "github.com/legocy-co/legocy/internal/domain/lego/repository"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"golang.org/x/net/context"
)

type LegoSeriesService struct {
	repo r.LegoSeriesRepository
}

func NewLegoSeriesService(repo r.LegoSeriesRepository) LegoSeriesService {
	return LegoSeriesService{repo: repo}
}

func (s *LegoSeriesService) ListSeries(ctx context.Context) ([]*models.LegoSeries, *errors.AppError) {
	return s.repo.GetLegoSeriesList(ctx)
}

func (s *LegoSeriesService) DetailSeries(ctx context.Context, id int) (*models.LegoSeries, *errors.AppError) {
	return s.repo.GetLegoSeries(ctx, id)
}

func (s *LegoSeriesService) CreateLegoSeries(ctx context.Context, m *models.LegoSeriesValueObject) *errors.AppError {
	return s.repo.CreateLegoSeries(ctx, m)
}

func (s *LegoSeriesService) DeleteSeries(ctx context.Context, id int) *errors.AppError {
	return s.repo.DeleteLegoSeries(ctx, id)
}

func (s *LegoSeriesService) UpdateSeries(legoSeriesID int, vo *models.LegoSeriesValueObject) *errors.AppError {
	return s.repo.UpdateLegoSeries(legoSeriesID, vo)
}

package v1

import (
	models "legocy-go/pkg/lego/models"
	repo "legocy-go/pkg/lego/repository"

	"golang.org/x/net/context"
)

type LegoSeriesService struct {
	repo repo.LegoSeriesRepository
}

func (s *LegoSeriesService) ListSeries(ctx context.Context) ([]*models.LegoSeries, error) {
	return s.repo.GetLegoSeries(ctx)
}

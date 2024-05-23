package service

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type MarketItemImageService struct {
	repo r.MarketItemImageRepository
}

func NewMarketItemImageService(repo r.MarketItemImageRepository) MarketItemImageService {
	return MarketItemImageService{
		repo: repo,
	}
}

func (s *MarketItemImageService) StoreMarketItemImage(
	vo models.MarketItemImageValueObject) (*models.MarketItemImage, *errors.AppError) {

	return s.repo.Store(vo)
}

func (s *MarketItemImageService) DeleteImageByID(id int) error {
	return s.repo.Delete(id)
}

func (s *MarketItemImageService) UpdateImageByID(id int, vo models.MarketItemImagePartialVO) (*models.MarketItemImage, *errors.AppError) {
	return s.repo.Update(id, vo)
}

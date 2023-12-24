package service

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
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

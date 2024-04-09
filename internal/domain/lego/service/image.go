package service

import (
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/domain/lego/repository"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type LegoSetImageService struct {
	repo repository.LegoSetImageRepository
}

func NewLegoSetImageService(repo repository.LegoSetImageRepository) LegoSetImageService {
	return LegoSetImageService{
		repo: repo,
	}
}

func (s *LegoSetImageService) GetLegoSetImages(legoSetID int) ([]*models.LegoSetImage, *errors.AppError) {
	return s.repo.Get(legoSetID)
}

func (s *LegoSetImageService) StoreLegoSetImage(vo models.LegoSetImageValueObject) (*models.LegoSetImage, *errors.AppError) {
	return s.repo.Store(vo)
}

func (s *LegoSetImageService) DeleteImageByID(id int) *errors.AppError {
	return s.repo.Delete(id)
}

package service

import (
	"context"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	repo "github.com/legocy-co/legocy/internal/domain/users/repository"
)

type UserImageService struct {
	repo repo.UserImageRepository
}

func NewUserImageUseCase(repo repo.UserImageRepository) UserImageService {
	return UserImageService{repo: repo}
}

func (s *UserImageService) StoreUserImage(c context.Context, image *models.UserImage) error {
	return s.repo.AddUserImage(c, image)
}

func (s *UserImageService) GetUserImages(c context.Context, userID int) ([]*models.UserImage, error) {
	return s.repo.GetUserImages(c, userID)
}

package auth

import (
	"context"
	repo "legocy-go/internal/db/postgres/repository"
	models "legocy-go/pkg/auth/models"
)

type UserImageUseCase struct {
	repo repo.UserImagePostgresRepository
}

func (s *UserImageUseCase) StoreUserImage(c context.Context, image *models.UserImage) error {
	return s.repo.AddUserImage(c, image)
}

func (s *UserImageUseCase) GetUserImages(c context.Context, userID int) ([]*models.UserImage, error) {
	return s.repo.GetUserImages(c, userID)
}

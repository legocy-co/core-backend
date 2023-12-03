package repository

import (
	"context"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserImageRepository interface {
	AddUserImage(c context.Context, image *models.UserImage) error
	GetUserImages(c context.Context, userID int) ([]*models.UserImage, error)
}

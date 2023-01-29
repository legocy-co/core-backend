package repository

import (
	"context"
	models "legocy-go/pkg/auth/models"
)

type UserImageRepository interface {
	AddUserImage(c context.Context, image *models.UserImage) error
	GetUserImages(c context.Context, userID int) ([]*models.UserImage, error)
}

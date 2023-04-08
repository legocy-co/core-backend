package postgres

import (
	models "legocy-go/internal/domain/users/models"
)

type UserPostgresImage struct {
	Model
	UserID      uint
	FilepathURL string
}

func (ui *UserPostgresImage) ToUserImage() *models.UserImage {
	return &models.UserImage{
		UserID:      int(ui.UserID),
		FilepathURL: ui.FilepathURL,
	}
}

func FromUserImage(ui *models.UserImage) *UserPostgresImage {
	return &UserPostgresImage{
		UserID:      uint(ui.UserID),
		FilepathURL: ui.FilepathURL,
	}
}

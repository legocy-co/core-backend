package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserPostgresImage struct {
	Model
	UserID      uint
	FilepathURL string
}

func (ui UserPostgresImage) TableName() string {
	return "user_images"
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

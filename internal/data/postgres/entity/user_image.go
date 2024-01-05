package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserImagePostgres struct {
	Model
	UserID      uint
	FilepathURL string
}

func (ui UserImagePostgres) TableName() string {
	return "user_images"
}

func (ui *UserImagePostgres) ToUserImage() *models.UserImage {
	return &models.UserImage{
		UserID:      int(ui.UserID),
		FilepathURL: ui.FilepathURL,
	}
}

func FromUserImage(ui *models.UserImage) *UserImagePostgres {
	return &UserImagePostgres{
		UserID:      uint(ui.UserID),
		FilepathURL: ui.FilepathURL,
	}
}

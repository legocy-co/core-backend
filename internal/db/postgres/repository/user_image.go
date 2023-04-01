package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	models "legocy-go/internal/domain/users/models"
	"legocy-go/internal/domain/users/repository"
)

type UserImagePostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserImagePostgresRepository(conn d.DataBaseConnection) repository.UserImageRepository {
	return UserImagePostgresRepository{conn: conn}
}

func (r UserImagePostgresRepository) AddUserImage(c context.Context, image *models.UserImage) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromUserImage(image)
	result := db.Create(&entity)
	return result.Error
}

func (r UserImagePostgresRepository) GetUserImages(c context.Context, userID int) ([]*models.UserImage, error) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var userImagesDB []*entities.UserPostgresImage
	db.Model(&entities.UserPostgresImage{}).Find(
		&userImagesDB, entities.UserPostgresImage{UserID: uint(userID)})

	if len(userImagesDB) == 0 {
		return nil, d.ErrItemNotFound
	}

	userImages := make([]*models.UserImage, 0, len(userImagesDB))
	for _, entity := range userImagesDB {
		userImages = append(userImages, entity.ToUserImage())
	}

	return userImages, nil
}

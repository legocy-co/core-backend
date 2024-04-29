package postgres

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/domain/users/repository"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/kafka"
	"github.com/legocy-co/legocy/pkg/kafka/schemas"
)

type UserImagePostgresRepository struct {
	conn       d.DBConn
	dispatcher events.Dispatcher
}

func NewUserImagePostgresRepository(conn d.DBConn, dispatcher events.Dispatcher) repository.UserImageRepository {
	return UserImagePostgresRepository{
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (r UserImagePostgresRepository) AddUserImage(c context.Context, image *models.UserImage) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	err := r.DeleteImagesByUserID(c, image.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	entity := entities.FromUserImage(image)
	result := tx.Create(&entity)

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()
	return nil
}

func (r UserImagePostgresRepository) DeleteImagesByUserID(c context.Context, userID int) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	var userImagesDB []*entities.UserImagePostgres
	db.Model(
		&entities.UserImagePostgres{},
	).Find(&userImagesDB, entities.UserImagePostgres{UserID: uint(userID)})

	for _, userImage := range userImagesDB {
		err := tx.Delete(&entities.UserImagePostgres{}, userImage.ID).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		err = r.dispatcher.ProduceJSONEvent(
			kafka.UserImagesDeletedTopic,
			schemas.ImageDeletedEventData{
				ImageFilepath: userImage.FilepathURL,
			},
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (r UserImagePostgresRepository) GetUserImages(c context.Context, userID int) ([]*models.UserImage, error) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var userImagesDB []*entities.UserImagePostgres
	db.Model(&entities.UserImagePostgres{}).Find(
		&userImagesDB, entities.UserImagePostgres{UserID: uint(userID)})

	if len(userImagesDB) == 0 {
		return nil, d.ErrItemNotFound
	}

	userImages := make([]*models.UserImage, 0, len(userImagesDB))
	for _, entity := range userImagesDB {
		userImages = append(userImages, entity.ToUserImage())
	}

	return userImages, nil
}

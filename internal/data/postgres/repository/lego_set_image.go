package postgres

import (
	"github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/kafka"
	"github.com/legocy-co/legocy/pkg/kafka/schemas"
)

type LegoSetImagePostgresRepository struct {
	conn       data.DBConn
	dispatcher events.Dispatcher
}

func NewLegoSetImagePostgresRepository(conn data.DBConn, dispatcher events.Dispatcher) LegoSetImagePostgresRepository {
	return LegoSetImagePostgresRepository{
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (r LegoSetImagePostgresRepository) Get(legoSetID int) ([]*models.LegoSetImage, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	var images []*entities.LegoSetImagePostgres
	if err := db.Where("lego_set_id = ?", legoSetID).Find(&images).Error; err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	var result []*models.LegoSetImage
	for _, image := range images {
		result = append(result, image.ToLegoSetImage())
	}

	return result, nil
}

func (r LegoSetImagePostgresRepository) Store(vo models.LegoSetImageValueObject) (*models.LegoSetImage, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	tx := db.Begin()

	image := entities.FromLegoSetImageValueObject(vo)
	if err := tx.Create(image).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	tx.Commit()
	return image.ToLegoSetImage(), nil
}

func (r LegoSetImagePostgresRepository) Delete(id int) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &data.ErrConnectionLost
	}

	tx := db.Begin()

	var image *entities.LegoSetImagePostgres
	if err := tx.Where("id = ?", id).Find(&image).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	if image == nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, "image not found")
		return &appErr
	}

	if err := tx.Delete(&entities.LegoSetImagePostgres{}, id).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	err := r.dispatcher.ProduceJSONEvent(
		kafka.LegoSetImagesDeletedTopic,
		schemas.ImageDeletedEventData{
			ImageFilepath: image.ImageURL,
		},
	)
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

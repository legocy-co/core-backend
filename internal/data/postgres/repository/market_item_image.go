package postgres

import (
	"github.com/legocy-co/legocy/internal/data"
	e "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/kafka"
	"github.com/legocy-co/legocy/pkg/kafka/schemas"
	log "github.com/sirupsen/logrus"
)

type MarketItemImagePostgresRepository struct {
	conn       data.DBConn
	dispatcher events.Dispatcher
}

func NewMarketItemImagePostgresRepository(conn data.DBConn, dispatcher events.Dispatcher) *MarketItemImagePostgresRepository {
	return &MarketItemImagePostgresRepository{
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (r MarketItemImagePostgresRepository) Store(
	vo models.MarketItemImageValueObject,
) (
	*models.MarketItemImage, *errors.AppError,
) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	tx := db.Begin()

	entityToCreate := e.FromMarketItemImageValueObject(vo)

	if err := tx.Create(entityToCreate).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	tx.Commit()
	return entityToCreate.ToMarketItemImage(), nil
}

func (r MarketItemImagePostgresRepository) Get(marketItemID int) ([]*models.MarketItemImage, *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	var marketItemImagesDB []*e.MarketItemImagePostgres

	err := db.Model(
		&e.MarketItemImagePostgres{},
	).Find(
		&marketItemImagesDB, e.MarketItemImagePostgres{MarketItemID: uint(marketItemID)},
	).Error

	if err != nil {
		appErr := errors.NewAppError(errors.NotFoundError, err.Error())
		return nil, &appErr
	}

	markItemImages := make([]*models.MarketItemImage, 0, len(marketItemImagesDB))
	for _, entity := range marketItemImagesDB {
		markItemImages = append(markItemImages, entity.ToMarketItemImage())
	}

	return markItemImages, nil
}

func (r MarketItemImagePostgresRepository) GetByID(id int) (*models.MarketItemImage, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	var marketItemImageDB *e.MarketItemImagePostgres

	err := db.Model(
		&e.MarketItemImagePostgres{},
	).First(
		&marketItemImageDB, id,
	).Error

	if err != nil {
		appErr := errors.NewAppError(errors.NotFoundError, err.Error())
		return nil, &appErr
	}

	return marketItemImageDB.ToMarketItemImage(), nil
}

func (r MarketItemImagePostgresRepository) Update(
	id int, vo models.MarketItemImagePartialVO,
) (*models.MarketItemImage, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &data.ErrConnectionLost
	}

	tx := db.Begin()

	err := tx.Exec("UPDATE market_item_images SET sort_index = ? WHERE id = ?", vo.SortIndex, id).Error
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	tx.Commit()
	return nil, nil
}

func (r MarketItemImagePostgresRepository) Delete(id int) error {
	db := r.conn.GetDB()

	if db == nil {
		return &data.ErrConnectionLost
	}

	tx := db.Begin()

	var currentImage *e.MarketItemImagePostgres
	if err := tx.First(&currentImage, id).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.NotFoundError, err.Error())
		return &appErr
	}

	if currentImage == nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.NotFoundError, "image not found")
		return &appErr
	}

	if err := tx.Delete(&e.MarketItemImagePostgres{}, id).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	err := r.dispatcher.ProduceJSONEvent(
		kafka.MarketItemImagesDeletedTopic,
		schemas.ImageDeletedEventData{
			ImageFilepath: currentImage.ImageURL,
		},
	)
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.InternalError,
			err.Error(),
		)
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r MarketItemImagePostgresRepository) DeleteByMarketItemId(marketItemId int) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &data.ErrConnectionLost
	}

	tx := db.Begin()

	var currentImages []*e.MarketItemImagePostgres
	if err := tx.Find(&currentImages, e.MarketItemImagePostgres{MarketItemID: uint(marketItemId)}).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.NotFoundError, err.Error())
		return &appErr
	}

	for _, image := range currentImages {
		err := r.Delete(int(image.ID))
		if err != nil {
			log.Printf("Error deleting image: %v", err.Error())
		}
	}

	tx.Commit()
	return nil
}

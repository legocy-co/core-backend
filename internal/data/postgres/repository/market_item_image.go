package postgres

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	"github.com/legocy-co/legocy/internal/data"
	e "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/kafka"
	"github.com/legocy-co/legocy/pkg/kafka/schemas"
)

type MarketItemImagePostgresRepository struct {
	conn data.DataBaseConnection
}

func NewMarketItemImagePostgresRepository(conn data.DataBaseConnection) *MarketItemImagePostgresRepository {
	return &MarketItemImagePostgresRepository{
		conn: conn,
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

	err := kafka.ProduceJSONEvent(
		kafka.MARKET_ITEM_IMAGES_DELETED_TOPIC,
		schemas.ImageDeletedEventData{
			ImageFilepath: currentImage.ImageURL,
		},
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

package admin

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/events"
)

type MarketItemAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemAdminPostgresRepository(conn d.DataBaseConnection) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{conn: conn}
}

func (m MarketItemAdminPostgresRepository) GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, *errors.AppError) {

	db := m.conn.GetDB()
	if db == nil {
		return []*models.MarketItemAdmin{}, &d.ErrConnectionLost
	}

	var itemsDB []*entities.MarketItemPostgres
	res := db.Model(
		&entities.MarketItemPostgres{}).
		Preload("Seller").
		Preload("LegoSet").
		Preload("LegoSet.LegoSeries").
		Preload("Currency").
		Preload("Location").
		Find(&itemsDB)

	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return []*models.MarketItemAdmin{}, &appErr
	}

	marketItemsAdmin := make([]*models.MarketItemAdmin, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItemsAdmin = append(marketItemsAdmin, entity.ToMarketItemAdmin())
	}

	return marketItemsAdmin, nil

}

func (m MarketItemAdminPostgresRepository) GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError) {
	db := m.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	query := db.Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&entity, "id = ? and status = 'ACTIVE'", id)

	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return nil, &appErr
	}

	// Not Found
	if query.RowsAffected == 0 {
		return nil, &e.ErrMarketItemsNotFound
	}

	return entity.ToMarketItemAdmin(), nil
}

func (m MarketItemAdminPostgresRepository) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError {
	db := m.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemAdminValueObject(*vo)
	if entity == nil {
		return &d.ErrItemNotFound
	}

	result := tx.Create(&entity)

	tx.Commit()

	err := events.ProduceJSONEvent(
		events.MARKET_ITEM_UPDATES_TOPIC,
		map[string]interface{}{
			"itemID": int(entity.ID),
		},
	)
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.InternalError, err.Error())
		return &appErr
	}

	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}

func (m MarketItemAdminPostgresRepository) UpdateMarketItemByID(
	c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError) {
	db := m.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, itemId)
	if entity == nil {
		return nil, &e.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItemAdmin(*vo)
	db.Save(entityUpdated)

	return m.GetMarketItemByID(c, itemId)
}

func (m MarketItemAdminPostgresRepository) DeleteMarketItemByID(c context.Context, itemId int) *errors.AppError {
	db := m.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	result := db.Delete(entities.MarketItemPostgres{}, itemId)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}

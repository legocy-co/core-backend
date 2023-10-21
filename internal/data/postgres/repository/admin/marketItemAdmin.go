package admin

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/pkg/kafka"
)

type MarketItemAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemAdminPostgresRepository(conn d.DataBaseConnection) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{conn: conn}
}

func (m MarketItemAdminPostgresRepository) GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, error) {

	db := m.conn.GetDB()
	if db == nil {
		return []*models.MarketItemAdmin{}, d.ErrConnectionLost
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
		return []*models.MarketItemAdmin{}, res.Error
	}

	marketItemsAdmin := make([]*models.MarketItemAdmin, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItemsAdmin = append(marketItemsAdmin, entity.ToMarketItemAdmin())
	}

	return marketItemsAdmin, nil

}

func (m MarketItemAdminPostgresRepository) GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, error) {
	db := m.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	result := db.Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&entity, "id = ? and status = 'ACTIVE'", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToMarketItemAdmin(), nil
}

func (m MarketItemAdminPostgresRepository) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) error {
	db := m.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemAdminValueObject(*vo)
	if entity == nil {
		return d.ErrItemNotFound
	}

	result := tx.Create(&entity)

	err := kafka.ProduceJSONEvent(
		kafka.MARKET_ITEM_UPDATES_TOPIC, map[string]interface{}{
			"itemID": int(entity.ID),
		})
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return result.Error
}

func (m MarketItemAdminPostgresRepository) UpdateMarketItemByID(
	c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, error) {
	db := m.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, itemId)
	if entity == nil {
		return nil, errors.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItemAdmin(*vo)
	db.Save(entityUpdated)

	return m.GetMarketItemByID(c, itemId)
}

func (m MarketItemAdminPostgresRepository) DeleteMarketItemByID(c context.Context, itemId int) error {
	db := m.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	result := db.Delete(entities.MarketItemPostgres{}, itemId)
	return result.Error
}

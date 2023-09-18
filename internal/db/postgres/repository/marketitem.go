package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entity"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/pkg/filter"
	"legocy-go/pkg/kafka"
)

type MarketItemPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemPostgresRepository(conn d.DataBaseConnection) MarketItemPostgresRepository {
	return MarketItemPostgresRepository{conn: conn}
}

func (r MarketItemPostgresRepository) GetMarketItems(
	c context.Context) ([]*models.MarketItem, error) {

	var itemsDB []*entities.MarketItemPostgres
	pagination := c.Value("pagination").(*filter.QueryParams)

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	res := db.Model(&entities.MarketItemPostgres{}).
		Scopes(filter.FilterDbByQueryParams(pagination, filter.PAGINATE)).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB, "status = 'ACTIVE'")
	if res.Error != nil {
		return nil, res.Error
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemsAuthorized(
	c context.Context, userID int) ([]*models.MarketItem, error) {

	var itemsDB []*entities.MarketItemPostgres
	pagination := c.Value("pagination").(*filter.QueryParams)

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	res := db.Model(&entities.MarketItemPostgres{}).
		Scopes(filter.FilterDbByQueryParams(pagination, filter.PAGINATE)).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB, "user_postgres_id <> ? and status = 'ACTIVE'", userID)
	if res.Error != nil {
		return nil, res.Error
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemByID(
	c context.Context, id int) (*models.MarketItem, error) {

	db := r.conn.GetDB()
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

	return entity.ToMarketItem(), nil
}

func (r MarketItemPostgresRepository) GetMarketItemsBySellerID(
	c context.Context, sellerID int) ([]*models.MarketItem, error) {

	var itemsDB []*entities.MarketItemPostgres
	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	result := db.Model(&entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB, "user_postgres_id = ? and status = 'ACTIVE'", sellerID)
	if result.Error != nil {
		return nil, result.Error
	}

	marketItems := make([]*models.MarketItem, 0, len(itemsDB))
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) GetMarketItemSellerID(
	c context.Context, id int) (int, error) {

	var count int

	db := r.conn.GetDB()
	if db == nil {
		return count, d.ErrConnectionLost
	}

	err := db.Model(entities.MarketItemPostgres{}).Where(
		"id=?", id).Select("user_postgres_id").First(&count).Error

	return count, err
}

func (r MarketItemPostgresRepository) GetSellerMarketItemsAmount(
	c context.Context, sellerID int) (int64, error) {

	var count int64

	db := r.conn.GetDB()
	if db == nil {
		return count, d.ErrConnectionLost
	}

	res := db.Model(
		entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).Count(&count)

	return count, res.Error
}

func (r MarketItemPostgresRepository) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) error {

	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.FromMarketItemValueObject(item)
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

	return result.Error
}

func (r MarketItemPostgresRepository) DeleteMarketItem(c context.Context, id int) error {

	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	result := db.Delete(entities.MarketItemPostgres{}, id)
	return result.Error
}

func (r MarketItemPostgresRepository) UpdateMarketItemByID(
	c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, error) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, id)
	if entity == nil {
		return nil, errors.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItem(*item)
	db.Save(entityUpdated)

	return r.GetMarketItemByID(c, id)
}

func (r MarketItemPostgresRepository) UpdateMarketItemByIDAdmin(
	c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, error) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.MarketItemPostgres
	_ = db.First(&entity, id)
	if entity == nil {
		return nil, errors.ErrMarketItemsNotFound
	}

	entityUpdated := entity.GetUpdatedMarketItemAdmin(*item)
	db.Save(entityUpdated)

	return r.GetMarketItemByID(c, id)
}

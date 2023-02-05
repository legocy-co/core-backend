package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	"legocy-go/pkg/filter"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemPostgresRepository(conn d.DataBaseConnection) MarketItemPostgresRepository {
	return MarketItemPostgresRepository{conn: conn}
}

func (r MarketItemPostgresRepository) GetMarketItems(
	c context.Context) ([]*models.MarketItem, error) {

	var marketItems []*models.MarketItem
	var itemsDB []*entities.MarketItemPostgres
	pagination := c.Value("pagination").(*filter.QueryParams)

	db := r.conn.GetDB()
	if db == nil {
		return marketItems, d.ErrConnectionLost
	}

	res := db.Model(&entities.MarketItemPostgres{}).
		Scopes(filter.FilterDbByQueryParams(pagination, filter.ALL)).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB)
	if res.Error != nil {
		return marketItems, res.Error
	}

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
	result := db.First(&entity, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity.ToMarketItem(), nil
}

func (r MarketItemPostgresRepository) GetSellerMarketItemsAmount(
	c context.Context, sellerID int) (int64, error) {

	var count *int64

	db := r.conn.GetDB()
	if db == nil {
		return *count, d.ErrConnectionLost
	}

	res := db.Where(
		entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).
		Count(count)

	return *count, res.Error
}

func (r MarketItemPostgresRepository) GetMarketItemsBySeller(
	c context.Context, sellerID int) ([]*models.MarketItem, error) {

	var marketItems []*models.MarketItem
	var itemsDB []*entities.MarketItemPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	err := db.Where(&entities.MarketItemPostgres{UserPostgresID: uint(sellerID)}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB)

	if err.Error != nil {
		return nil, err.Error
	}

	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) CreateMarketItem(
	c context.Context, item *models.MarketItemBasic) error {

	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromMarketItemBasic(item)
	if entity == nil {
		return d.ErrItemNotFound
	}

	result := db.Create(&entity)
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

package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entities"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemPostgresRepository(conn d.DataBaseConnection) MarketItemPostgresRepository {
	return MarketItemPostgresRepository{conn: conn}
}

func (r MarketItemPostgresRepository) GetMarketItems(c context.Context) ([]*models.MarketItem, error) {
	var marketItems []*models.MarketItem

	db := r.conn.GetDB()
	if db == nil {
		return marketItems, d.ErrConnectionLost
	}

	var itemsDB []*entities.MarketItemPostgres
	res := db.Model(&entities.MarketItemPostgres{}).
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

func (r MarketItemPostgresRepository) GetMarketItemByID(c context.Context, id int) (*models.MarketItem, error) {

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

func (r MarketItemPostgresRepository) GetMarketItemsBySeller(
	c context.Context, sellerID int) ([]*models.MarketItem, error) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var itemsDB []*entities.MarketItemPostgres
	err := db.Model(&entities.MarketItemPostgres{}).
		Preload("Seller").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").Preload("Location").
		Find(&itemsDB, entities.MarketItemPostgres{UserPostgresID: uint(sellerID)})

	if err.Error != nil {
		return nil, err.Error
	}

	var marketItems []*models.MarketItem
	for _, entity := range itemsDB {
		marketItems = append(marketItems, entity.ToMarketItem())
	}

	return marketItems, nil
}

func (r MarketItemPostgresRepository) CreateMarketItem(c context.Context,
	item *models.MarketItemBasic) error {

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

package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entity"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemAdminPostgresRepository(conn d.DataBaseConnection) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{conn: conn}
}

func (r MarketItemAdminPostgresRepository) GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, error) {

	db := r.conn.GetDB()
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

func (m MarketItemAdminPostgresRepository) GetMarketItemByID(c context.Context) (*models.MarketItemAdmin, error) {
	//TODO implement me
	panic("implement me")
}

func (m MarketItemAdminPostgresRepository) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) error {
	//TODO implement me
	panic("implement me")
}

func (m MarketItemAdminPostgresRepository) UpdateMarketItemByID(c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, error) {
	//TODO implement me
	panic("implement me")
}

func (m MarketItemAdminPostgresRepository) DeleteMarketItemByID(c context.Context, itemId int) error {
	//TODO implement me
	panic("implement me")
}

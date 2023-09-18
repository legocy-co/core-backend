package postgres

import (
	"context"
	d "legocy-go/internal/db"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewMarketItemAdminPostgresRepository(conn d.DataBaseConnection) MarketItemAdminPostgresRepository {
	return MarketItemAdminPostgresRepository{conn: conn}
}

func (m MarketItemAdminPostgresRepository) GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, error) {
	//TODO implement me
	panic("implement me")
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

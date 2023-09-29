package admin

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
	marketplace "legocy-go/internal/domain/marketplace/repository"
)

type MarketItemAdminService struct {
	repo marketplace.MarketItemAdminRepository
}

func NewMarketItemAdminService(r marketplace.MarketItemAdminRepository) MarketItemAdminService {
	return MarketItemAdminService{repo: r}
}

func (s MarketItemAdminService) GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, error) {
	return s.repo.GetMarketItems(c)
}

func (s MarketItemAdminService) GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, error) {
	return s.repo.GetMarketItemByID(c, id)
}

func (s MarketItemAdminService) CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) error {
	return s.repo.CreateMarketItem(c, vo)
}

func (s MarketItemAdminService) UpdateMarketItem(c context.Context, id int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, error) {
	return s.repo.UpdateMarketItemByID(c, id, vo)
}

func (s MarketItemAdminService) DeleteMarketItemById(c context.Context, id int) error {
	return s.repo.DeleteMarketItemByID(c, id)
}

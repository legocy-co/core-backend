package marketplace

import (
	"golang.org/x/net/context"
	models "legocy-go/internal/domain/marketplace/models"
	r "legocy-go/internal/domain/marketplace/repository"
)

type MarketItemService struct {
	repo r.MarketItemRepository
}

func NewMarketItemSerivce(repo r.MarketItemRepository) MarketItemService {
	return MarketItemService{repo: repo}
}

func (ms *MarketItemService) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) error {
	return ms.repo.CreateMarketItem(c, item)
}

func (ms *MarketItemService) ListMarketItems(
	c context.Context) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItems(c)
}

func (ms *MarketItemService) MarketItemDetail(c context.Context, id int) (*models.MarketItem, error) {
	return ms.repo.GetMarketItemByID(c, id)
}

func (ms *MarketItemService) DeleteMarketItem(c context.Context, id int) error {
	return ms.repo.DeleteMarketItem(c, id)
}

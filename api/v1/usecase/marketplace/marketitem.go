package marketplace

import (
	models "legocy-go/pkg/marketplace/models"
	r "legocy-go/pkg/marketplace/repository"

	"golang.org/x/net/context"
)

type MarketItemService struct {
	repo r.MarketItemRepository
}

func NewMarketItemSerivce(repo r.MarketItemRepository) MarketItemService {
	return MarketItemService{repo: repo}
}

func (ms *MarketItemService) CreateMarketItem(
	c context.Context, item *models.MarketItemBasic) error {
	return ms.repo.CreateMarketItem(c, item)
}

func (ms *MarketItemService) ListMarketItems(
	c context.Context) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItems(c)
}

func (ms *MarketItemService) SellerMarketItems(
	c context.Context, id int) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItemsBySeller(c, id)
}

func (ms *MarketItemService) MarketItemDetail(c context.Context, id int) (*models.MarketItem, error) {
	return ms.repo.GetMarketItemByID(c, id)
}

func (ms *MarketItemService) DeleteMarketItem(c context.Context, id int) error {
	return ms.repo.DeleteMarketItem(c, id)
}

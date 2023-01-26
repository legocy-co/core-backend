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

func (ms *MarketItemService) CreateMarketItem(c context.Context, item *models.MarketItem) error {
	return ms.repo.CreateMarketItem(c, item)
}

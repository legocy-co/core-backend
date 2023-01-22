package marketplace

import (
	"context"
	auth "legocy-go/pkg/auth/models"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemRepository interface {
	GetMarketItems(c context.Context) ([]*models.MarketItem, error)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, error)
	GetMarketItemsBySeller(c context.Context, seller auth.User) ([]*models.MarketItem, error)
	CreateMarketItem(c context.Context, item *models.MarketItem) error
}

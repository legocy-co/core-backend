package marketplace

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemRepository interface {
	GetMarketItems(c context.Context) ([]*models.MarketItem, error)
	GetMarketItemsAuthorized(c context.Context, userID int) ([]*models.MarketItem, error)
	GetMarketItemsBySellerID(c context.Context, sellerID int) ([]*models.MarketItem, error)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, error)
	GetMarketItemSellerID(c context.Context, id int) (int, error)
	GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, error)
	CreateMarketItem(c context.Context, item *models.MarketItemValueObject) error
	UpdateMarketItemByID(
		c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, error)
	UpdateMarketItemByIDAdmin(c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, error)
	DeleteMarketItem(c context.Context, id int) error
}

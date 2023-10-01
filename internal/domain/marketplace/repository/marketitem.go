package marketplace

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
	admin "legocy-go/internal/domain/marketplace/models/admin"
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
	UpdateMarketItemByIDAdmin(
		c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, error)
	DeleteMarketItem(c context.Context, id int) error
}

type MarketItemAdminRepository interface {
	GetMarketItems(c context.Context) ([]*admin.MarketItemAdmin, error)
	GetMarketItemByID(c context.Context, id int) (*admin.MarketItemAdmin, error)
	CreateMarketItem(c context.Context, vo *admin.MarketItemAdminValueObject) error
	UpdateMarketItemByID(c context.Context, itemId int, vo *admin.MarketItemAdminValueObject) (*admin.MarketItemAdmin, error)
	DeleteMarketItemByID(c context.Context, itemId int) error
}

package marketplace

import (
	"context"
	"legocy-go/internal/app/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemRepository interface {
	GetMarketItems(c context.Context) ([]*models.MarketItem, *errors.AppError)
	GetMarketItemsAuthorized(c context.Context, userID int) ([]*models.MarketItem, *errors.AppError)
	GetMarketItemsBySellerID(c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError)
	GetMarketItemSellerID(c context.Context, id int) (int, *errors.AppError)
	GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, *errors.AppError)
	CreateMarketItem(c context.Context, item *models.MarketItemValueObject) *errors.AppError
	UpdateMarketItemByID(
		c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError)
	DeleteMarketItem(c context.Context, id int) *errors.AppError
}

type MarketItemAdminRepository interface {
	GetMarketItems(c context.Context) ([]*models.MarketItemAdmin, *errors.AppError)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError)
	CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError
	UpdateMarketItemByID(
		c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError)
	DeleteMarketItemByID(c context.Context, itemId int) *errors.AppError
}

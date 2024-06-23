package marketplace

import (
	"context"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/lib/pagination"
)

type MarketItemRepository interface {

	// GetMarketItems is a Public Method, which doesn't exclude current user items
	GetMarketItems(
		ctx pagination.PaginationContext,
		filter *domain.MarketItemFilterCriteria) (pagination.Page[*models.MarketItem], *errors.AppError)

	// GetMarketItemsAuthorized is an Authorized Method, which excludes current user items
	GetMarketItemsAuthorized(
		ctx pagination.PaginationContext,
		filter *domain.MarketItemFilterCriteria,
		userID int) (pagination.Page[*models.MarketItem], *errors.AppError)

	// GetActiveMarketItemsBySellerID gets all active market items for a specific user
	GetActiveMarketItemsBySellerID(
		c context.Context,
		sellerID int) ([]*models.MarketItem, *errors.AppError)

	// GetMarketItemsBySellerID gets all market items for a specific user
	GetMarketItemsBySellerID(c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError)

	// GetMarketItemByID gets a market item by its ID
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError)

	// GetActiveMarketItemByID gets a market item by its ID, only if it's active
	GetActiveMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError)

	GetMarketItemSellerID(c context.Context, id int) (int, *errors.AppError)

	GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, *errors.AppError)

	CreateMarketItem(c context.Context, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError)

	UpdateMarketItemByID(c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError)

	DeleteMarketItem(c context.Context, id int) *errors.AppError
}

type MarketItemAdminRepository interface {
	GetMarketItems(ctx pagination.PaginationContext) (pagination.Page[*models.MarketItemAdmin], *errors.AppError)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError)
	CreateMarketItem(c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError
	UpdateMarketItemByID(
		c context.Context, itemId int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError)
	DeleteMarketItemByID(c context.Context, itemId int) *errors.AppError
}

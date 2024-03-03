package marketplace

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/pkg/pagination"
)

type MarketItemRepository interface {
	GetMarketItems(ctx pagination.PaginationContext, filter *domain.MarketItemFilterCriteria) (pagination.Page[*models.MarketItem], *errors.AppError)
	GetMarketItemsAuthorized(ctx pagination.PaginationContext, filter *domain.MarketItemFilterCriteria, userID int) (pagination.Page[*models.MarketItem], *errors.AppError)
	GetMarketItemsBySellerID(c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError)
	GetMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError)
	GetMarketItemSellerID(c context.Context, id int) (int, *errors.AppError)
	GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, *errors.AppError)
	GetPendingMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError)
	CreateMarketItem(c context.Context, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError)
	UpdateMarketItemByID(
		c context.Context, id int, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError)
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

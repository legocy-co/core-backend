package service

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	"github.com/legocy-co/legocy/pkg/pagination"
	"golang.org/x/net/context"
)

type MarketItemService struct {
	repo r.MarketItemRepository
}

func NewMarketItemService(repo r.MarketItemRepository) MarketItemService {
	return MarketItemService{repo: repo}
}

func (ms *MarketItemService) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError) {
	return ms.repo.CreateMarketItem(c, item)
}

func (ms *MarketItemService) ListMarketItems(
	c pagination.PaginationContext) (pagination.Page[*models.MarketItem], *errors.AppError) {

	marketItems, err := ms.repo.GetMarketItems(c)
	if err != nil {
		return marketItems, err
	}

	if len(marketItems.GetObjects()) == 0 {
		return marketItems, &e.ErrMarketItemsNotFound
	}

	return marketItems, err
}

func (ms *MarketItemService) ListMarketItemsAuthorized(
	c pagination.PaginationContext, userID int) (pagination.Page[*models.MarketItem], *errors.AppError) {

	marketItems, err := ms.repo.GetMarketItemsAuthorized(c, userID)
	if err != nil {
		return marketItems, err
	}

	if len(marketItems.GetObjects()) == 0 {
		return marketItems, &e.ErrMarketItemsNotFound
	}

	return marketItems, err
}

func (ms *MarketItemService) MarketItemsBySellerID(
	c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError) {
	return ms.repo.GetMarketItemsBySellerID(c, sellerID)
}

func (ms *MarketItemService) MarketItemDetail(
	c context.Context, id int) (*models.MarketItem, *errors.AppError) {
	return ms.repo.GetMarketItemByID(c, id)
}

func (ms *MarketItemService) DeleteMarketItem(c context.Context, id int) *errors.AppError {
	return ms.repo.DeleteMarketItem(c, id)
}

func (ms *MarketItemService) UpdateMarketItemByID(
	c context.Context, currentUserID int, id int, vo *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError) {

	if currentUserID != vo.SellerID {
		return nil, &e.ErrMarketItemInvalidSellerID
	}

	return ms.repo.UpdateMarketItemByID(c, id, vo)
}

func (ms *MarketItemService) GetMarketItemSellerID(c context.Context, id int) (int, *errors.AppError) {
	return ms.repo.GetMarketItemSellerID(c, id)
}

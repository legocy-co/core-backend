package service

import (
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/pkg/pagination"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type MarketItemService struct {
	imageRepo r.MarketItemImageRepository
	repo      r.MarketItemRepository
}

func NewMarketItemService(repo r.MarketItemRepository, imageRepo r.MarketItemImageRepository) MarketItemService {
	return MarketItemService{repo: repo, imageRepo: imageRepo}
}

func (ms *MarketItemService) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) (*models.MarketItem, *errors.AppError) {
	return ms.repo.CreateMarketItem(c, item)
}

func (ms *MarketItemService) ListMarketItems(
	c pagination.PaginationContext,
	filter *domain.MarketItemFilterCriteria,
) (pagination.Page[*models.MarketItem], *errors.AppError) {

	marketItems, err := ms.repo.GetMarketItems(c, filter)
	if err != nil {
		return marketItems, err
	}

	if len(marketItems.GetObjects()) == 0 {
		return marketItems, &e.ErrMarketItemsNotFound
	}

	return marketItems, err
}

func (ms *MarketItemService) ListMarketItemsAuthorized(
	c pagination.PaginationContext,
	filter *domain.MarketItemFilterCriteria,
	userID int,
) (pagination.Page[*models.MarketItem], *errors.AppError) {

	marketItems, err := ms.repo.GetMarketItemsAuthorized(c, filter, userID)
	if err != nil {
		return marketItems, err
	}

	if len(marketItems.GetObjects()) == 0 {
		return marketItems, &e.ErrMarketItemsNotFound
	}

	return marketItems, err
}

func (ms *MarketItemService) ActiveMarketItemsBySellerID(
	c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError) {
	return ms.repo.GetActiveMarketItemsBySellerID(c, sellerID)
}

func (ms *MarketItemService) ActiveMarketItemDetail(
	c context.Context, id int) (*models.MarketItem, *errors.AppError) {
	return ms.repo.GetActiveMarketItemByID(c, id)
}

func (ms *MarketItemService) DeleteMarketItem(c context.Context, id int) *errors.AppError {
	err := ms.imageRepo.DeleteByMarketItemId(id)
	if err != nil {
		log.Printf("Error deleting market item images: %v", err)
	}

	return ms.repo.DeleteMarketItem(c, id)
}

func (ms *MarketItemService) UpdateMarketItemByID(
	c context.Context,
	currentUserID int,
	id int,
	vo *models.MarketItemValueObject,
) (*models.MarketItem, *errors.AppError) {

	if currentUserID != vo.SellerID {
		return nil, &e.ErrMarketItemInvalidSellerID
	}

	return ms.repo.UpdateMarketItemByID(c, id, vo)
}

func (ms *MarketItemService) GetMarketItemSellerID(c context.Context, id int) (int, *errors.AppError) {
	return ms.repo.GetMarketItemSellerID(c, id)
}

func (ms *MarketItemService) GetSellerMarketItemsAmount(c context.Context, sellerID int) (int64, *errors.AppError) {
	return ms.repo.GetSellerMarketItemsAmount(c, sellerID)
}

func (ms *MarketItemService) GetMarketItemsBySellerID(c context.Context, sellerID int) ([]*models.MarketItem, *errors.AppError) {
	return ms.repo.GetMarketItemsBySellerID(c, sellerID)
}

func (ms *MarketItemService) GetMarketItemByID(c context.Context, id int) (*models.MarketItem, *errors.AppError) {
	return ms.repo.GetMarketItemByID(c, id)
}

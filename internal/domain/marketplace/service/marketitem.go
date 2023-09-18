package service

import (
	"golang.org/x/net/context"
	"legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	r "legocy-go/internal/domain/marketplace/repository"
)

type MarketItemService struct {
	repo r.MarketItemRepository
}

func NewMarketItemService(repo r.MarketItemRepository) MarketItemService {
	return MarketItemService{repo: repo}
}

func (ms *MarketItemService) CreateMarketItem(
	c context.Context, item *models.MarketItemValueObject) error {
	return ms.repo.CreateMarketItem(c, item)
}

func (ms *MarketItemService) ListMarketItems(
	c context.Context) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItems(c)
}

func (ms *MarketItemService) ListMarketItemsAuthorized(
	c context.Context, userID int) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItemsAuthorized(c, userID)
}

func (ms *MarketItemService) MarketItemsBySellerID(
	c context.Context, sellerID int) ([]*models.MarketItem, error) {
	return ms.repo.GetMarketItemsBySellerID(c, sellerID)
}

func (ms *MarketItemService) MarketItemDetail(
	c context.Context, id int) (*models.MarketItem, error) {
	return ms.repo.GetMarketItemByID(c, id)
}

func (ms *MarketItemService) DeleteMarketItem(c context.Context, id int) error {
	return ms.repo.DeleteMarketItem(c, id)
}

func (ms *MarketItemService) UpdateMarketItemByID(
	c context.Context, currentUserID int, id int, vo *models.MarketItemValueObject) (*models.MarketItem, error) {

	if currentUserID != vo.SellerID {
		return nil, errors.ErrMarketItemInvalidSellerID
	}

	return ms.repo.UpdateMarketItemByID(c, id, vo)
}

func (ms *MarketItemService) UpdateMarketItemByIDAdmin(
	c context.Context, id int, vo *models.MarketItemValueObject) (*models.MarketItem, error) {
	return ms.repo.UpdateMarketItemByIDAdmin(c, id, vo)
}

func (ms *MarketItemService) GetMarketItemSellerID(c context.Context, id int) (int, error) {
	return ms.repo.GetMarketItemSellerID(c, id)
}

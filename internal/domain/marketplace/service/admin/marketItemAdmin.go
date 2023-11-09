package admin

import (
	"context"
	"legocy-go/internal/domain/errors"
	models "legocy-go/internal/domain/marketplace/models"
	marketplace "legocy-go/internal/domain/marketplace/repository"
)

type MarketItemAdminService struct {
	repo marketplace.MarketItemAdminRepository
}

func NewMarketItemAdminService(
	r marketplace.MarketItemAdminRepository) MarketItemAdminService {
	return MarketItemAdminService{repo: r}
}

func (s MarketItemAdminService) GetMarketItems(
	c context.Context) ([]*models.MarketItemAdmin, *errors.AppError) {
	return s.repo.GetMarketItems(c)
}

func (s MarketItemAdminService) GetMarketItemByID(
	c context.Context, id int) (*models.MarketItemAdmin, *errors.AppError) {
	return s.repo.GetMarketItemByID(c, id)
}

func (s MarketItemAdminService) CreateMarketItem(
	c context.Context, vo *models.MarketItemAdminValueObject) *errors.AppError {
	return s.repo.CreateMarketItem(c, vo)
}

func (s MarketItemAdminService) UpdateMarketItem(
	c context.Context, id int, vo *models.MarketItemAdminValueObject) (*models.MarketItemAdmin, *errors.AppError) {
	return s.repo.UpdateMarketItemByID(c, id, vo)
}

func (s MarketItemAdminService) DeleteMarketItemById(c context.Context, id int) *errors.AppError {
	return s.repo.DeleteMarketItemByID(c, id)
}

package admin

import (
	"context"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	marketplace "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/pkg/pagination"
)

type MarketItemAdminService struct {
	repo marketplace.MarketItemAdminRepository
}

func NewMarketItemAdminService(
	r marketplace.MarketItemAdminRepository) MarketItemAdminService {
	return MarketItemAdminService{repo: r}
}

func (s MarketItemAdminService) GetMarketItems(
	ctx pagination.PaginationContext,
) (pagination.Page[*models.MarketItemAdmin], *errors.AppError) {
	return s.repo.GetMarketItems(ctx)
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

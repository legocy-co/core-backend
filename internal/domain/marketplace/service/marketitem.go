package service

import (
	e "github.com/legocy-co/legocy/internal/domain/marketplace/errors"
	domain "github.com/legocy-co/legocy/internal/domain/marketplace/filters"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	r "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/lib/pagination"
	"golang.org/x/net/context"
)

type MarketItemService struct {
	imageRepo r.MarketItemImageRepository
	likesRepo r.LikeRepository
	repo      r.MarketItemRepository
}

type MarketItemsServiceOpts struct {
	Repo      r.MarketItemRepository
	ImageRepo r.MarketItemImageRepository
	LikesRepo r.LikeRepository
}

func NewMarketItemService(opts MarketItemsServiceOpts) MarketItemService {
	return MarketItemService{
		repo:      opts.Repo,
		imageRepo: opts.ImageRepo,
		likesRepo: opts.LikesRepo,
	}
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
	if err := ms.imageRepo.DeleteByMarketItemId(id); err != nil {
		return err
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

func (ms *MarketItemService) GetLikedItems(
	c pagination.PaginationContext,
	userID int,
) (pagination.Page[*models.MarketItem], *errors.AppError) {

	// get liked items
	likedItems, err := ms.likesRepo.GetLikesByUserID(userID)
	if err != nil {
		return pagination.NewEmptyPage[*models.MarketItem](), err
	}

	// Get ids to filter
	var ids = make([]int, 0, len(likedItems))
	for _, item := range likedItems {
		ids = append(ids, item.MarketItemID())
	}

	filterCriteria := domain.MarketItemFilterCriteria{Ids: ids}

	return ms.repo.GetMarketItemsAuthorized(c, &filterCriteria, userID)
}

package service

import (
	"context"
	"github.com/legocy-co/legocy/internal/domain/lego"
	"github.com/legocy-co/legocy/internal/domain/lego/filters"
	models "github.com/legocy-co/legocy/internal/domain/lego/models"
	r "github.com/legocy-co/legocy/internal/domain/lego/repository"
	"github.com/legocy-co/legocy/internal/pkg/errors"
	"github.com/legocy-co/legocy/lib/pagination"
)

type LegoSetService struct {
	repo r.LegoSetRepository
}

func NewLegoSetService(repo r.LegoSetRepository) LegoSetService {
	return LegoSetService{repo: repo}
}

func (u *LegoSetService) ListLegoSets(c context.Context) ([]*models.LegoSet, *errors.AppError) {
	legoSets, err := u.repo.GetLegoSets(c)

	if err != nil {
		return legoSets, err
	}

	if len(legoSets) == 0 {
		return legoSets, &lego.ErrLegoSetsNotFound
	}

	return legoSets, nil
}

func (u *LegoSetService) GetSetsPage(ctx pagination.PaginationContext, filter *filters.LegoSetFilterCriteria) (pagination.Page[models.LegoSet], *errors.AppError) {
	return u.repo.GetSetsPage(ctx, filter)
}

func (u *LegoSetService) LegoSetDetail(c context.Context, id int) (*models.LegoSet, *errors.AppError) {
	return u.repo.GetLegoSetByID(c, id)
}

func (u *LegoSetService) LegoSetCreate(c context.Context, legoSet *models.LegoSetValueObject) (*models.LegoSet, *errors.AppError) {
	return u.repo.CreateLegoSet(c, legoSet)
}

func (u *LegoSetService) LegoSetDelete(c context.Context, id int) *errors.AppError {
	return u.repo.DeleteLegoSet(c, id)
}

func (u *LegoSetService) LegoSetUpdate(legoSetID int, vo *models.LegoSetValueObject) *errors.AppError {
	return u.repo.UpdateLegoSetByID(legoSetID, vo)
}

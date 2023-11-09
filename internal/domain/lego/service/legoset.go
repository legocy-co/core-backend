package service

import (
	"context"
	"legocy-go/internal/domain/errors"
	"legocy-go/internal/domain/lego"
	models "legocy-go/internal/domain/lego/models"
	r "legocy-go/internal/domain/lego/repository"
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

func (u *LegoSetService) LegoSetDetail(c context.Context, id int) (*models.LegoSet, *errors.AppError) {
	return u.repo.GetLegoSetByID(c, id)
}

func (u *LegoSetService) LegoSetCreate(c context.Context, legoSet *models.LegoSetValueObject) *errors.AppError {
	return u.repo.CreateLegoSet(c, legoSet)
}

func (u *LegoSetService) LegoSetDelete(c context.Context, id int) *errors.AppError {
	return u.repo.DeleteLegoSet(c, id)
}

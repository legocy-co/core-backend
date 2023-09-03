package service

import (
	"context"
	models "legocy-go/internal/domain/lego/models"
	r "legocy-go/internal/domain/lego/repository"
)

type LegoSetUseCase struct {
	repo r.LegoSetRepository
}

func NewLegoSetUseCase(repo r.LegoSetRepository) LegoSetUseCase {
	return LegoSetUseCase{repo: repo}
}

func (u *LegoSetUseCase) ListLegoSets(c context.Context) ([]*models.LegoSet, error) {
	return u.repo.GetLegoSets(c)
}

func (u *LegoSetUseCase) LegoSetDetail(c context.Context, id int) (*models.LegoSet, error) {
	return u.repo.GetLegoSetByID(c, id)
}

func (u *LegoSetUseCase) LegoSetCreate(c context.Context, legoSet *models.LegoSetValueObject) error {
	return u.repo.CreateLegoSet(c, legoSet)
}

func (u *LegoSetUseCase) LegoSetDelete(c context.Context, id int) error {
	return u.repo.DeleteLegoSet(c, id)
}

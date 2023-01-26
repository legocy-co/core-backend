package lego

import (
	"context"
	models "legocy-go/pkg/lego/models"
	r "legocy-go/pkg/lego/repository"
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

package service

import (
	"context"
	models "legocy-go/internal/domain/lego/models"
	r "legocy-go/internal/domain/lego/repository"
)

type LegoSetService struct {
	repo r.LegoSetRepository
}

func NewLegoSetService(repo r.LegoSetRepository) LegoSetService {
	return LegoSetService{repo: repo}
}

func (u *LegoSetService) ListLegoSets(c context.Context) ([]*models.LegoSet, error) {
	return u.repo.GetLegoSets(c)
}

func (u *LegoSetService) LegoSetDetail(c context.Context, id int) (*models.LegoSet, error) {
	return u.repo.GetLegoSetByID(c, id)
}

func (u *LegoSetService) LegoSetCreate(c context.Context, legoSet *models.LegoSetValueObject) error {
	return u.repo.CreateLegoSet(c, legoSet)
}

func (u *LegoSetService) LegoSetDelete(c context.Context, id int) error {
	return u.repo.DeleteLegoSet(c, id)
}

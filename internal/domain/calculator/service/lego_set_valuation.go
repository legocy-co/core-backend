package service

import (
	"context"
	"legocy-go/internal/app/errors"
	"legocy-go/internal/domain/calculator/models"
	"legocy-go/internal/domain/calculator/repository"
)

type LegoSetValuationService struct {
	r repository.LegoSetValuationRepository
}

func NewLegoSetValuationService(r repository.LegoSetValuationRepository) LegoSetValuationService {
	return LegoSetValuationService{r: r}
}

func (s LegoSetValuationService) GetLegoSetValuations(c context.Context, legoSetId int) ([]models.LegoSetValuation, *errors.AppError) {
	return s.r.GetLegoSetValuationsList(c, legoSetId)
}

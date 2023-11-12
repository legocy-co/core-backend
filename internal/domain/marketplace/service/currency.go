package service

import (
	"context"
	"legocy-go/internal/app/errors"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/internal/domain/marketplace/repository"
)

type CurrencyService struct {
	repo marketplace.CurrencyRepository
}

func NewCurrencyService(repo marketplace.CurrencyRepository) CurrencyService {
	return CurrencyService{repo: repo}
}

func (s CurrencyService) CurrenciesList(c context.Context) ([]*models.Currency, *errors.AppError) {
	return s.repo.GetCurrencies(c)
}

func (s CurrencyService) CurrencyDetail(c context.Context, symbol string) (*models.Currency, *errors.AppError) {
	return s.repo.GetCurrency(c, symbol)
}

func (s CurrencyService) CreateCurrency(c context.Context, curr *models.CurrencyValueObject) *errors.AppError {
	return s.repo.CreateCurrency(c, curr)
}

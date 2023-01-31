package marketplace

import (
	"context"
	models "legocy-go/pkg/marketplace/models"
	marketplace "legocy-go/pkg/marketplace/repository"
)

type CurrencyUseCase struct {
	repo marketplace.CurrencyRepository
}

func NewCurrencyUseCase(repo marketplace.CurrencyRepository) CurrencyUseCase {
	return CurrencyUseCase{repo: repo}
}

func (s *CurrencyUseCase) CurrenciesList(c context.Context) ([]*models.Currency, error) {
	return s.repo.GetCurrencies(c)
}

package marketplace

import marketplace "legocy-go/pkg/marketplace/repository"

type CurrencyUseCase struct {
	repo marketplace.CurrencyRepository
}

func NewCurrencyUseCase(repo marketplace.CurrencyRepository) CurrencyUseCase {
	return CurrencyUseCase{repo: repo}
}

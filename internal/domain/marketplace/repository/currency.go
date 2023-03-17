package marketplace

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
)

type CurrencyRepository interface {
	GetCurrencies(c context.Context) ([]*models.Currency, error)
	GetCurrency(c context.Context, symbol string) (*models.Currency, error)
	CreateCurrency(c context.Context, currency *models.CurrencyBasic) error
}

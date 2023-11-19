package marketplace

import (
	"context"
	"legocy-go/internal/app/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type CurrencyRepository interface {
	GetCurrencies(c context.Context) ([]*models.Currency, *errors.AppError)
	GetCurrency(c context.Context, symbol string) (*models.Currency, *errors.AppError)
	CreateCurrency(c context.Context, currency *models.CurrencyValueObject) *errors.AppError
}

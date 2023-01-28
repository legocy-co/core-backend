package postgres

import (
	"legocy-go/internal/db/postgres"
	models "legocy-go/pkg/marketplace/models"
)

type CurrencyPostgres struct {
	postgres.Model
	Name   string `gorm:"unique"`
	Symbol string `gorm:"unique"`
}

func (cp *CurrencyPostgres) ToCurrency() *models.Currency {
	return &models.Currency{
		ID:     int(cp.ID),
		Name:   cp.Name,
		Symbol: cp.Symbol,
	}
}

func FromCurrency(curr *models.Currency) *CurrencyPostgres {
	return &CurrencyPostgres{
		Name:   curr.Name,
		Symbol: curr.Symbol,
	}
}

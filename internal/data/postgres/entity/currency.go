package postgres

import (
	models "legocy-go/internal/domain/marketplace/models"
)

type CurrencyPostgres struct {
	Model
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

func FromCurrencyValueObject(curr *models.CurrencyValueObject) *CurrencyPostgres {
	return &CurrencyPostgres{
		Name:   curr.Name,
		Symbol: curr.Symbol,
	}
}

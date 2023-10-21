package models

import (
	lego "legocy-go/internal/domain/lego/models"
	marketplace "legocy-go/internal/domain/marketplace/models"
)

type LegoSetValuation struct {
	ID               int
	LegoSet          lego.LegoSet
	State            string
	CompanyValuation float32 // legocy price
	Currency         marketplace.Currency
}

type LegoSetValuationValueObject struct {
	LegoSetID        int
	State            string
	CompanyValuation float32
	CurrencyID       marketplace.Currency
}

package models

import (
	lego "github.com/legocy-co/legocy/internal/domain/lego/models"
)

type LegoSetValuation struct {
	ID               int
	LegoSet          lego.LegoSet
	State            string
	CompanyValuation float32 // legocy price
}

type LegoSetValuationValueObject struct {
	LegoSetID        int
	State            string
	CompanyValuation float32
}

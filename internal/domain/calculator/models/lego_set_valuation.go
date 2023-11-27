package models

import (
	lego "legocy-go/internal/domain/lego/models"
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

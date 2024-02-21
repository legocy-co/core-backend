package models

import calculator "github.com/legocy-co/legocy/internal/domain/calculator/models"

type SetWithValuation struct {
	CollectionSet CollectionLegoSet
	SetValuation  *calculator.LegoSetValuation
}

func NewSetWithValuation(set CollectionLegoSet, valuation *calculator.LegoSetValuation) SetWithValuation {
	return SetWithValuation{
		CollectionSet: set,
		SetValuation:  valuation,
	}
}

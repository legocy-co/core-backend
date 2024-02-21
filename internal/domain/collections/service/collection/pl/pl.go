package pl

import (
	"github.com/legocy-co/legocy/internal/domain/collections/models"
)

func GetCollectionSetProfits(collectionSet models.SetWithValuation) *models.CollectionLegoSetProfits {

	if collectionSet.SetValuation == nil {
		return &models.CollectionLegoSetProfits{
			ReturnPercentage: 0,
			ReturnUSD:        0,
		}
	}

	set := collectionSet.CollectionSet
	valuation := collectionSet.SetValuation

	gr := ((valuation.CompanyValuation - set.BuyPrice) / set.BuyPrice) * 100
	value := valuation.CompanyValuation - set.BuyPrice

	return &models.CollectionLegoSetProfits{
		ReturnPercentage: gr,
		ReturnUSD:        value,
	}
}

func GetCollectionProfits(collectionSets []models.SetWithValuation) *models.CollectionProfits {

	var totalBuyPrice float32
	var totalCurrentValuation float32

	for _, collectionSet := range collectionSets {
		totalBuyPrice += collectionSet.CollectionSet.BuyPrice
		if collectionSet.SetValuation != nil {
			totalCurrentValuation += collectionSet.SetValuation.CompanyValuation
		}
	}

	totalReturnUSD := totalCurrentValuation - totalBuyPrice
	totalReturnPercent := (totalReturnUSD / totalBuyPrice) * 100

	return &models.CollectionProfits{
		TotalReturnUSD:        totalReturnUSD,
		TotalReturnPercentage: totalReturnPercent,
	}
}

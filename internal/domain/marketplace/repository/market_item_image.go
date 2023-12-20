package marketplace

import models "github.com/legocy-co/legocy/internal/domain/marketplace/models"

type MarketItemImageRepository interface {
	Store(vo models.MarketItemImageValueObject) (*models.MarketItemImage, error)
	Delete(id int) error
}

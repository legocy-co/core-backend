package marketplace

import (
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type MarketItemImageRepository interface {
	Get(marketItemID int) ([]*models.MarketItemImage, *errors.AppError)
	Store(vo models.MarketItemImageValueObject) (*models.MarketItemImage, *errors.AppError)
	Delete(id int) error
}

package marketplace

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type MarketItemImageRepository interface {
	Get(marketItemID int) ([]*models.MarketItemImage, *errors.AppError)
	GetByID(id int) (*models.MarketItemImage, *errors.AppError)
	Store(vo models.MarketItemImageValueObject) (*models.MarketItemImage, *errors.AppError)
	Update(id int, vo models.MarketItemImagePartialVO) (*models.MarketItemImage, *errors.AppError)
	Delete(id int) error
	DeleteByMarketItemId(marketItemId int) *errors.AppError
}

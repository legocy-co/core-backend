package postgres

import models "github.com/legocy-co/legocy/internal/domain/marketplace/models"

type MarketItemImagePostgres struct {
	Model
	MarketItemID uint               `gorm:"not null"`
	MarketItem   MarketItemPostgres `gorm:"foreignKey:MarketItemID"`
	ImageURL     string             `gorm:"not null"`
	IsMain       bool               `gorm:"default=false"`
}

func (m *MarketItemImagePostgres) TableName() string {
	return "market_item_images"
}

func FromMarketItemImageValueObject(vo models.MarketItemImageValueObject) *MarketItemImagePostgres {
	return &MarketItemImagePostgres{
		MarketItemID: uint(vo.MarketItemID),
		ImageURL:     vo.ImageURL,
		IsMain:       vo.IsMain,
	}
}

func (m *MarketItemImagePostgres) ToMarketItemImage() *models.MarketItemImage {
	return &models.MarketItemImage{
		MarketItemID: int(m.MarketItemID),
		ImageURL:     m.ImageURL,
		IsMain:       m.IsMain,
	}
}

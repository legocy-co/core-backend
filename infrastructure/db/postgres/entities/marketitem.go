package postgres

import (
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemPostgres struct {
	Model
	Price              float32
	CurrencyPostgresID uint
	Currency           CurrencyPostgres `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LegoSetPostgresID  uint
	LegoSet            LegoSetPostgres `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserPostgresID     uint
	Seller             UserPostgres `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	LocationPostgresID uint
	Location           LocationPostgres `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (mp *MarketItemPostgres) ToMarketItem() *models.MarketItem {
	return &models.MarketItem{
		LegoSet:  *mp.LegoSet.ToLegoSet(),
		Seller:   *mp.Seller.ToUser(),
		Price:    mp.Price,
		Currency: *mp.Currency.ToCurrency(),
		Location: *mp.Location.ToLocation(),
	}
}

func FromMarketItem(mi *models.MarketItem) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:              mi.Price,
		CurrencyPostgresID: uint(mi.Currency.ID),
		LegoSetPostgresID:  uint(mi.LegoSet.ID),
		UserPostgresID:     uint(mi.Seller.ID),
		LocationPostgresID: uint(mi.Location.ID),
	}
}

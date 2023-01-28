package postgres

import (
	"legocy-go/infrastructure/db/postgres"
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemPostgres struct {
	postgres.Model
	Price              float32
	CurrencyPostgresID uint
	Currency           CurrencyPostgres `gorm:"ForeignKey:CurrencyPostgresID"`
	LegoSetPostgresID  uint
	LegoSet            LegoSetPostgres `gorm:"ForeignKey:LegoSetPostgresID"`
	UserPostgresID     uint
	Seller             UserPostgres `gorm:"ForeignKey:UserPostgresID"`
	LocationPostgresID uint
	Location           LocationPostgres `gorm:"ForeignKey:LocationPostgresID"`
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

package postgres

import (
	models "legocy-go/pkg/marketplace/models"
)

type MarketItemPostgres struct {
	Model
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
		ID:       int(mp.ID),
		LegoSet:  *mp.LegoSet.ToLegoSet(),
		Seller:   *mp.Seller.ToUser(),
		Price:    mp.Price,
		Currency: *mp.Currency.ToCurrency(),
		Location: *mp.Location.ToLocation(),
	}
}

func FromMarketItemBasic(mi *models.MarketItemBasic) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:              mi.Price,
		CurrencyPostgresID: uint(mi.CurrencyID),
		LegoSetPostgresID:  uint(mi.LegoSetID),
		UserPostgresID:     uint(mi.SellerID),
		LocationPostgresID: uint(mi.LocationID),
	}
}

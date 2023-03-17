package postgres

import (
	models "legocy-go/internal/domain/marketplace/models"
)

type MarketItemPostgres struct {
	Model
	Price              float32
	CurrencyPostgresID uint             `filter:"param:currencyId;searchable,filterable"`
	Currency           CurrencyPostgres `gorm:"ForeignKey:CurrencyPostgresID"`
	LegoSetPostgresID  uint             `filter:"param:setId;searchable,filterable"`
	LegoSet            LegoSetPostgres  `gorm:"ForeignKey:LegoSetPostgresID"`
	UserPostgresID     uint             `filter:"param:sellerId;searchable,filterable"`
	Seller             UserPostgres     `gorm:"ForeignKey:UserPostgresID"`
	LocationPostgresID uint             `filter:"param:locationId;searchable,filterable"`
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

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

func FromMarketItemValueObject(mi *models.MarketItemValueObject) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:              mi.Price,
		CurrencyPostgresID: uint(mi.CurrencyID),
		LegoSetPostgresID:  uint(mi.LegoSetID),
		UserPostgresID:     uint(mi.SellerID),
		LocationPostgresID: uint(mi.LocationID),
	}
}

func (mp *MarketItemPostgres) GetUpdatedMarketItem(vo models.MarketItemValueObject) *MarketItemPostgres {
	mp.CurrencyPostgresID = uint(vo.CurrencyID)
	mp.LegoSetPostgresID = uint(vo.LegoSetID)
	mp.LocationPostgresID = uint(vo.LocationID)
	mp.Price = vo.Price
	mp.UserPostgresID = uint(vo.SellerID)

	return mp
}

func FromMarketItemAdminValueObject(vo models.MarketItemAdminValueObject) *MarketItemPostgres {
	return &MarketItemPostgres{
		Price:              vo.Price,
		CurrencyPostgresID: uint(vo.CurrencyID),
		LegoSetPostgresID:  uint(vo.LegoSetID),
		UserPostgresID:     uint(vo.SellerID),
		LocationPostgresID: uint(vo.LocationID),
	}
}

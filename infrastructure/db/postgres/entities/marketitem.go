package postgres

type LocationPostgres struct {
	Model
	Country string `gorm:"uniqueIndex:idx_country_city"`
	City    string `gorm:"uniqueIndex:idx_country_city""`
}

type CurrencyPostgres struct {
	Model
	Name   string `gorm:"unique"`
	Symbol string `gorm:"unique"`
}

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

package postgres

type MarketItemImagePostgres struct {
	Model
	MarketItemID uint               `gorm:"not null"`
	MarketItem   MarketItemPostgres `gorm:"foreignKey:MarketItemID"`
	ImageURL     string             `gorm:"not null"`
}

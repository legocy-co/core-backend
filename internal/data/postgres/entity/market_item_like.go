package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"time"
)

type MarketItemLikePostgres struct {
	UserID int           `gorm:"primaryKey"`
	User   *UserPostgres `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	MarketItemID int                 `gorm:"primaryKey"`
	MarketItem   *MarketItemPostgres `gorm:"foreignKey:MarketItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	LikeDate time.Time `gorm:"not null" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}

func (l *MarketItemLikePostgres) TableName() string {
	return "market_item_likes"
}

func (l *MarketItemLikePostgres) ToMarketItemLike() *models.Like {
	return models.NewLike(l.MarketItemID, l.UserID)
}

func (l *MarketItemLikePostgres) ToMarketItemLikeAggregate() *models.Like {
	return models.NewLikeAggregate(l.MarketItemID, l.MarketItem.ToMarketItem(), l.UserID, l.User.ToUser())
}

func (l *MarketItemLikePostgres) UpdatedFields(vo *models.LikeValueObject) {
	l.UserID = vo.UserID
	l.MarketItemID = vo.MarketItemID
}

func FromMarketItemLikeValueObject(vo *models.LikeValueObject) *MarketItemLikePostgres {
	return &MarketItemLikePostgres{
		UserID:       vo.UserID,
		MarketItemID: vo.MarketItemID,
	}
}

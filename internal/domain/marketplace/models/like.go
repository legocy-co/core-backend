package marketplace

import users "github.com/legocy-co/legocy/internal/domain/users/models"

type Like struct {
	marketItemID int
	marketItem   *MarketItem
	userID       int
	user         *users.User
}

func NewLike(marketItemID, userID int) *Like {
	return &Like{
		marketItemID: marketItemID,
		userID:       userID,
	}
}

func NewLikeAggregate(marketItemID int, marketItem *MarketItem, userID int, user *users.User) *Like {
	return &Like{
		marketItemID: marketItemID,
		marketItem:   marketItem,
		user:         user,
		userID:       userID,
	}
}

func (l *Like) MarketItemID() int {
	return l.marketItemID
}

func (l *Like) MarketItem() *MarketItem {
	return l.marketItem
}

func (l *Like) UserID() int {
	return l.userID
}

func (l *Like) User() *users.User {
	return l.user
}

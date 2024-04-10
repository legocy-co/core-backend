package marketplace

import models "github.com/legocy-co/legocy/internal/domain/marketplace/models"

type LikeResponse struct {
	MarketItemID int `json:"market_item_id"`
	UserID       int `json:"user_id"`
}

func FromLikeDomain(like *models.Like) LikeResponse {
	return LikeResponse{
		MarketItemID: like.MarketItemID(),
		UserID:       like.UserID(),
	}
}

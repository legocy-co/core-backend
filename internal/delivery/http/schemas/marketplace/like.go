package marketplace

import models "github.com/legocy-co/legocy/internal/domain/marketplace/models"

type LikeResponse struct {
	MarketItemID int `json:"marketItemID"`
	UserID       int `json:"userID"`
}

func FromLikeDomain(like *models.Like) LikeResponse {
	return LikeResponse{
		MarketItemID: like.MarketItemID(),
		UserID:       like.UserID(),
	}
}

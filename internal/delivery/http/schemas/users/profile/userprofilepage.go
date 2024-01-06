package profile

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
)

type UserProfilePageResponse struct {
	User        users.UserDetailResponse         `json:"user"`
	UserReviews []users.UserReviewResponse       `json:"user_reviews"`
	MarketItems []marketplace.MarketItemResponse `json:"market_items"`
}

func GetUserProfilePageResponse(
	marketItems []marketplace.MarketItemResponse,
	user users.UserDetailResponse,
	userReviews []users.UserReviewResponse,
) UserProfilePageResponse {
	return UserProfilePageResponse{
		MarketItems: marketItems,
		User:        user,
		UserReviews: userReviews,
	}
}

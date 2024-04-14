package profile

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
)

type UserProfilePageResponse struct {
	User        users.UserDetailResponse         `json:"user"`
	UserReviews []users.UserReviewResponse       `json:"userReviews"`
	MarketItems []marketplace.MarketItemResponse `json:"marketItems"`
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

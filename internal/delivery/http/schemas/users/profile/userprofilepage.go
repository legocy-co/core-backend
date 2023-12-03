package profile

import (
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/marketplace"
	"github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
)

type UserProfilePageResponse struct {
	MarketItems []marketplace.MarketItemResponse `json:"marketItems"`
	User        users.UserDetailResponse         `json:"user"`
	UserReviews []users.UserReviewResponse       `json:"userReviews"`
	UserImages  []users.UserImageInfoResponse    `json:"userImages"`
}

func GetUserProfilePageResponse(
	marketItems []marketplace.MarketItemResponse, user users.UserDetailResponse,
	userReviews []users.UserReviewResponse, userImages []users.UserImageInfoResponse) UserProfilePageResponse {
	return UserProfilePageResponse{
		MarketItems: marketItems,
		User:        user,
		UserReviews: userReviews,
		UserImages:  userImages,
	}
}

package profile

import (
	"legocy-go/internal/delievery/http/resources/marketplace"
	"legocy-go/internal/delievery/http/resources/users"
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

package marketplace

import (
	"legocy-go/internal/delievery/http/resources/users"
)

type UserProfilePageResponse struct {
	MarketItems []MarketItemResponse          `json:"marketItems"`
	User        users.UserDetailResponse      `json:"user"`
	UserReviews []UserReviewResponse          `json:"userReviews"`
	UserImages  []users.UserImageInfoResponse `json:"userImages"`
}

func GetUserProfilePageResponse(
	marketItems []MarketItemResponse, user users.UserDetailResponse,
	userReviews []UserReviewResponse, userImages []users.UserImageInfoResponse) UserProfilePageResponse {
	return UserProfilePageResponse{
		MarketItems: marketItems,
		User:        user,
		UserReviews: userReviews,
		UserImages:  userImages,
	}
}

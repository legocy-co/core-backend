package users

import "legocy-go/internal/delievery/http/resources/marketplace"

type UserProfilePageResponse struct {
	MarketItems []marketplace.MarketItemResponse `json:"marketItems"`
	User        UserDetailResponse               `json:"user"`
	UserReviews []UserReviewResponse             `json:"userReviews"`
	UserImages  []UserImageInfoResponse          `json:"userImages"`
}

func GetUserProfilePageResponse(
	marketItems []marketplace.MarketItemResponse, user UserDetailResponse,
	userReviews []UserReviewResponse, userImages []UserImageInfoResponse) UserProfilePageResponse {
	return UserProfilePageResponse{
		MarketItems: marketItems,
		User:        user,
		UserReviews: userReviews,
		UserImages:  userImages,
	}
}

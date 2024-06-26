package users

import (
	"time"

	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type UserReviewRequest struct {
	SellerID int    `json:"sellerID"`
	Rating   int    `json:"rating"`
	Message  string `json:"message"`
}

func (r *UserReviewRequest) ToUserReviewValueObject(reviewerID int) (*models.UserReviewValueObject, error) {
	return models.NewUserReviewValueObject(
		r.SellerID,
		reviewerID,
		r.Rating,
		r.Message,
		time.Now().Format("02.01.06"),
	)
}

type UserReviewResponse struct {
	ID       int                `json:"id"`
	Rating   int                `json:"rating"`
	Message  string             `json:"message"`
	Date     string             `json:"date"`
	Seller   UserDetailResponse `json:"seller"`
	Reviewer UserDetailResponse `json:"reviewer"`
}

func GetUserReviewResponse(m *models.UserReview) UserReviewResponse {
	return UserReviewResponse{
		ID:       m.ID,
		Rating:   m.Rating,
		Message:  m.Message,
		Date:     m.Date,
		Seller:   GetUserDetailResponse(&m.Seller),
		Reviewer: GetUserDetailResponse(&m.Reviewer),
	}
}

type UserReviewTotalsResponse struct {
	AvgRating    float64 `json:"avgRating"`
	TotalReviews int     `json:"totalReviews"`
}

func GetUserReviewsTotalsResponse(m *models.UserRevewTotals) *UserReviewTotalsResponse {
	return &UserReviewTotalsResponse{
		AvgRating:    m.AvgRating,
		TotalReviews: m.TotalReviews,
	}
}

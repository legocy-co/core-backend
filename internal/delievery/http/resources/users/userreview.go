package users

import (
	models "legocy-go/internal/domain/marketplace/models"
	"time"
)

type UserReviewRequest struct {
	SellerID int    `json:"seller_id"`
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

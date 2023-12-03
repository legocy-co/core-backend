package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
)

type UserReviewPostgres struct {
	Model
	Rating             int
	Message            string
	SellerPostgresID   uint         `filter:"param:sellerId; searchable, filterable"`
	Seller             UserPostgres `gorm:"ForeignKey:SellerPostgresID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ReviewerPostgresID uint         `filter:"param:reviewerId; searchable, filterable"`
	Reviewer           UserPostgres `gorm:"ForeignKey:ReviewerPostgresID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date               string
}

func (urp UserReviewPostgres) TableName() string {
	return "user_reviews"
}

func (urp *UserReviewPostgres) ToUserReview() (*models.UserReview, error) {
	return models.NewUserReview(
		int(urp.ID),
		*urp.Seller.ToUser(),
		*urp.Reviewer.ToUser(),
		urp.Rating,
		urp.Message,
		urp.Date,
	)
}

func FromUserReviewValueObject(rev *models.UserReviewValueObject) *UserReviewPostgres {
	return &UserReviewPostgres{
		Rating:             rev.Rating,
		Message:            rev.Message,
		SellerPostgresID:   uint(rev.SellerID),
		ReviewerPostgresID: uint(rev.ReviewerID),
		Date:               rev.Date,
	}
}

package postgres

import (
	models "legocy-go/internal/domain/marketplace/models"
)

type UserReviewPostgres struct {
	Model
	Rating             int
	Message            string
	SellerPostgresID   uint         `filter:"param:sellerId; searchable, filterable" gorm:"primaryKey"`
	Seller             UserPostgres `gorm:"ForeignKey:SellerPostgresID"`
	ReviewerPostgresID uint         `filter:"param:reviewerId; searchable, filterable" gorm:"primaryKey"`
	Reviewer           UserPostgres `gorm:"ForeignKey:ReviewerPostgresID"`
	Date               string
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

package postgres

import (
	models "legocy-go/internal/domain/marketplace/models"
	"time"
)

type UserReviewPostgres struct {
	Model
	Rating             int
	Message            string
	SellerPostgresID   uint         `filter:"param:sellerId; searchable, filterable" gorm:"uniqueIndex:compositeindex"`
	Seller             UserPostgres `gorm:"ForeignKey:SellerPostgresID" gorm:"uniqueIndex:compositeindex"`
	ReviewerPostgresID uint         `filter:"param:reviewerId; searchable, filterable"`
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
		time.Now().Format("02.01.06"),
	)
}

func FromUserReviewValueObject(rev *models.UserReviewValueObject) *UserReviewPostgres {
	return &UserReviewPostgres{
		Rating:             rev.Rating,
		Message:            rev.Message,
		SellerPostgresID:   uint(rev.SellerID),
		ReviewerPostgresID: uint(rev.ReviewerID),
		Date:               time.Now().Format("02.01.06"),
	}
}

/*func (urp *UserReviewPostgres) ToUserReview() *models.UserReview {
	return &models.UserReview{
		ID:       int(urp.ID),
		Seller:   *urp.Seller.ToUser(),
		Reviewer: *urp.Reviewer.ToUser(),
		Rating:   urp.Rating,
		Message:  urp.Message,
		Date:     time.Now().Format("02.01.06"),
	}
}*/

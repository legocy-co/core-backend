package marketplace

import (
	"context"
	"legocy-go/internal/domain/errors"
	models "legocy-go/internal/domain/marketplace/models"
)

type UserReviewRepository interface {
	GetUserReviews(c context.Context) ([]*models.UserReview, *errors.AppError)
	GetUserReviewsBySellerID(c context.Context, sellerID int) ([]*models.UserReview, *errors.AppError)
	GetUserReviewByID(c context.Context, id int) (*models.UserReview, *errors.AppError)
	GetReviewerID(c context.Context, id int) (int, *errors.AppError)
	CreateUserReview(c context.Context, review *models.UserReviewValueObject) *errors.AppError
	DeleteUserReview(c context.Context, id int) *errors.AppError
}

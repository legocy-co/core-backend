package marketplace

import (
	"context"

	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type UserReviewRepository interface {
	GetUserReviews(c context.Context) ([]*models.UserReview, *errors.AppError)
	GetUserReviewsTotals(c context.Context, sellerID int) (*models.UserRevewTotals, *errors.AppError)
	GetUserReviewsBySellerID(c context.Context, sellerID int) ([]*models.UserReview, *errors.AppError)
	GetUserReviewByID(c context.Context, id int) (*models.UserReview, *errors.AppError)
	GetReviewerID(c context.Context, id int) (int, *errors.AppError)
	CreateUserReview(c context.Context, review *models.UserReviewValueObject) *errors.AppError
	DeleteUserReview(c context.Context, id int) *errors.AppError
}

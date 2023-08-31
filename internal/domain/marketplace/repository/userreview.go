package marketplace

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
)

type UserReviewRepository interface {
	GetUserReviews(c context.Context) ([]*models.UserReview, error)
	GetUserReviewByID(c context.Context, id int) (*models.UserReview, error)
	GetReviewerID(c context.Context, id int) (int, error)
	CreateUserReview(c context.Context, review *models.UserReviewValueObject) error
	DeleteUserReview(c context.Context, id int) error
}

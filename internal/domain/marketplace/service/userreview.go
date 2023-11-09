package service

import (
	"golang.org/x/net/context"
	"legocy-go/internal/domain/errors"
	e "legocy-go/internal/domain/marketplace/errors"
	models "legocy-go/internal/domain/marketplace/models"
	r "legocy-go/internal/domain/marketplace/repository"
)

type UserReviewService struct {
	repo r.UserReviewRepository
}

func NewUserReviewService(repo r.UserReviewRepository) UserReviewService {
	return UserReviewService{repo: repo}
}

func (ms *UserReviewService) CreateUserReview(
	c context.Context, review *models.UserReviewValueObject) *errors.AppError {
	return ms.repo.CreateUserReview(c, review)
}

func (ms *UserReviewService) ListUserReviews(
	c context.Context) ([]*models.UserReview, *errors.AppError) {
	userReviews, err := ms.repo.GetUserReviews(c)
	if err != nil {
		return userReviews, err
	}

	if len(userReviews) == 0 {
		return userReviews, &e.ErrUserReviewsNotFound
	}

	return userReviews, nil
}

func (ms *UserReviewService) UserReviewsBySellerID(
	c context.Context, sellerID int) ([]*models.UserReview, *errors.AppError) {
	return ms.repo.GetUserReviewsBySellerID(c, sellerID)
}

func (ms *UserReviewService) UserReviewDetail(c context.Context, id int) (*models.UserReview, *errors.AppError) {
	return ms.repo.GetUserReviewByID(c, id)
}

func (ms *UserReviewService) DeleteUserReview(c context.Context, id int) *errors.AppError {
	return ms.repo.DeleteUserReview(c, id)
}

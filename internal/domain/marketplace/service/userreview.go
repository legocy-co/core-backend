package service

import (
	"golang.org/x/net/context"
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
	c context.Context, review *models.UserReviewValueObject) error {
	return ms.repo.CreateUserReview(c, review)
}

func (ms *UserReviewService) ListUserReviews(
	c context.Context) ([]*models.UserReview, error) {
	return ms.repo.GetUserReviews(c)
}

func (ms *UserReviewService) UserReviewsBySellerID(
	c context.Context, sellerID int) ([]*models.UserReview, error) {
	return ms.repo.GetUserReviewsBySellerID(c, sellerID)
}

func (ms *UserReviewService) UserReviewDetail(c context.Context, id int) (*models.UserReview, error) {
	return ms.repo.GetUserReviewByID(c, id)
}

func (ms *UserReviewService) DeleteUserReview(c context.Context, id int) error {
	return ms.repo.DeleteUserReview(c, id)
}

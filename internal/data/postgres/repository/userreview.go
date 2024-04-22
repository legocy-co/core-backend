package postgres

import (
	"context"

	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
)

type UserReviewPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserReviewPostgresRepository(
	conn d.DataBaseConnection) UserReviewPostgresRepository {
	return UserReviewPostgresRepository{conn: conn}
}

func (r *UserPostgresRepository) GetUserReviewsTotals(c context.Context, sellerID int) (*models.UserRevewTotals, *errors.AppError) {

	db := r.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var totalReviews int64
	var averageRating float64

	query := db.Model(&entities.UserReviewPostgres{SellerPostgresID: uint(sellerID)})

	if err := query.Count(&totalReviews).Error; err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	if err := query.Select("AVG(rating) as average_rating").Scan(&averageRating).Error; err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return nil, &appErr
	}

	return models.NewUserRevewTotals(averageRating, int(totalReviews)), nil
}

func (r *UserReviewPostgresRepository) GetUserReviews(
	c context.Context) ([]*models.UserReview, *errors.AppError) {

	var itemsDB []*entities.UserReviewPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	res := db.Model(&entities.UserReviewPostgres{}).
		Preload("Reviewer").
		Preload("Seller").
		Find(&itemsDB)
	if res.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &appErr
	}

	userReviews := make([]*models.UserReview, 0, len(itemsDB))
	for _, entity := range itemsDB {
		userReview, err := entity.ToUserReview()
		if err != nil {
			appErr := errors.NewAppError(errors.InternalError, err.Error())
			return nil, &appErr
		}

		userReviews = append(userReviews, userReview)
	}

	return userReviews, nil
}

func (r *UserReviewPostgresRepository) GetUserReviewByID(
	c context.Context, id int) (*models.UserReview, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.UserReviewPostgres
	result := db.Preload("Reviewer").
		Preload("Seller").
		First(&entity, id)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return nil, &appErr
	}

	userReview, err := entity.ToUserReview()
	if err != nil {
		appErr := errors.NewAppError(errors.InternalError, err.Error())
		return userReview, &appErr
	}

	return userReview, nil
}

func (r *UserReviewPostgresRepository) GetUserReviewsBySellerID(
	c context.Context, sellerID int) ([]*models.UserReview, *errors.AppError) {

	var userReviewsDB []*entities.UserReviewPostgres
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	result := db.Model(&entities.UserReviewPostgres{SellerPostgresID: uint(sellerID)}).
		Preload("Reviewer").
		Preload("Seller").
		Find(&userReviewsDB, "seller_postgres_id = ?", sellerID)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return nil, &appErr
	}

	if len(userReviewsDB) == 0 {
		return nil, &d.ErrItemNotFound
	}

	userReviews := make([]*models.UserReview, 0, len(userReviewsDB))
	for _, entity := range userReviewsDB {
		userReview, err := entity.ToUserReview()

		if err != nil {
			return userReviews, nil
		}

		userReviews = append(userReviews, userReview)
	}

	return userReviews, nil
}

func (r *UserReviewPostgresRepository) GetReviewerID(
	c context.Context, id int) (int, *errors.AppError) {

	var count int

	db := r.conn.GetDB()
	if db == nil {
		return count, &d.ErrConnectionLost
	}

	err := db.Model(entities.UserReviewPostgres{}).
		Where("id = ?", id).Select("user_postgres_id").First(&count).Error

	var appErr *errors.AppError
	if err != nil {
		*appErr = errors.NewAppError(errors.ConflictError, err.Error())
	}

	return count, appErr
}

func (r *UserReviewPostgresRepository) CreateUserReview(
	c context.Context, review *models.UserReviewValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	entity := entities.FromUserReviewValueObject(review)
	if entity == nil {
		return &d.ErrItemNotFound
	}

	result := db.Create(&entity)

	var appErr *errors.AppError
	if result.Error != nil {
		*appErr = errors.NewAppError(errors.ConflictError, result.Error.Error())
	}

	return appErr
}

func (r *UserReviewPostgresRepository) DeleteUserReview(c context.Context, id int) *errors.AppError {

	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	result := db.Delete(entities.UserReviewPostgres{}, id)

	var appErr *errors.AppError
	if result.Error != nil {
		*appErr = errors.NewAppError(errors.ConflictError, result.Error.Error())
	}

	return appErr
}

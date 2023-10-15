package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/pkg/filter"
)

type UserReviewPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserReviewPostgresRepository(
	conn d.DataBaseConnection) UserReviewPostgresRepository {
	return UserReviewPostgresRepository{conn: conn}
}

func (r UserReviewPostgresRepository) GetUserReviews(
	c context.Context) ([]*models.UserReview, error) {

	var itemsDB []*entities.UserReviewPostgres
	pagination := c.Value("pagination").(*filter.QueryParams)

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	res := db.Model(&entities.UserReviewPostgres{}).
		Scopes(filter.FilterDbByQueryParams(pagination, filter.PAGINATE)).
		Preload("Reviewer").
		Preload("Seller").
		Find(&itemsDB)
	if res.Error != nil {
		return nil, res.Error
	}

	userReviews := make([]*models.UserReview, 0, len(itemsDB))
	for _, entity := range itemsDB {
		userReview, err := entity.ToUserReview()
		if err != nil {
			return nil, err
		}

		userReviews = append(userReviews, userReview)
	}

	return userReviews, nil
}

func (r UserReviewPostgresRepository) GetUserReviewByID(
	c context.Context, id int) (*models.UserReview, error) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var entity *entities.UserReviewPostgres
	result := db.Preload("Reviewer").
		Preload("Seller").
		First(&entity, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return entity.ToUserReview()
}

func (r UserReviewPostgresRepository) GetUserReviewsBySellerID(
	c context.Context, sellerID int) ([]*models.UserReview, error) {

	var userReviewsDB []*entities.UserReviewPostgres
	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	result := db.Model(&entities.UserReviewPostgres{SellerPostgresID: uint(sellerID)}).
		Preload("Reviewer").
		Preload("Seller").
		Find(&userReviewsDB, "seller_postgres_id = ?", sellerID)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(userReviewsDB) == 0 {
		return nil, d.ErrItemNotFound
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

func (r UserReviewPostgresRepository) GetReviewerID(
	c context.Context, id int) (int, error) {

	var count int

	db := r.conn.GetDB()
	if db == nil {
		return count, d.ErrConnectionLost
	}

	err := db.Model(entities.UserReviewPostgres{}).
		Where("id = ?", id).Select("user_postgres_id").First(&count).Error

	return count, err
}

func (r UserReviewPostgresRepository) CreateUserReview(
	c context.Context, review *models.UserReviewValueObject) error {

	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	entity := entities.FromUserReviewValueObject(review)
	if entity == nil {
		return d.ErrItemNotFound
	}

	result := db.Create(&entity)
	return result.Error
}

func (r UserReviewPostgresRepository) DeleteUserReview(c context.Context, id int) error {

	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	result := db.Delete(entities.UserReviewPostgres{}, id)
	return result.Error
}

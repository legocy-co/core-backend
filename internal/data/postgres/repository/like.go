package postgres

import (
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	models "github.com/legocy-co/legocy/internal/domain/marketplace/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

type LikePostgresRepository struct {
	conn d.Storage
}

func NewLikePostgresRepository(conn d.Storage) LikePostgresRepository {
	return LikePostgresRepository{conn: conn}
}

func (r LikePostgresRepository) AddLike(vo models.LikeValueObject) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &postgres.ErrConnectionLost
	}

	tx := db.Begin()

	likeDB := entities.FromMarketItemLikeValueObject(&vo)

	if err := tx.Create(likeDB).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r LikePostgresRepository) RemoveLike(vo models.LikeValueObject) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &postgres.ErrConnectionLost
	}

	tx := db.Begin()

	err := tx.Delete(
		entities.MarketItemLikePostgres{},
		"market_item_id = ? AND user_id = ?",
		vo.MarketItemID, vo.UserID,
	).Error

	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r LikePostgresRepository) GetLikesByUserID(userID int) ([]*models.Like, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	var likes []entities.MarketItemLikePostgres

	query := db.Model(
		&entities.MarketItemLikePostgres{},
	).Find(
		&likes, "user_id = ?", userID,
	)
	if query.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, query.Error.Error())
		return []*models.Like{}, &appErr
	}

	if query.RowsAffected == 0 {
		appErr := errors.NewAppError(errors.NotFoundError, "No likes found")
		return []*models.Like{}, &appErr
	}

	likesDomain := make([]*models.Like, 0, len(likes))
	for _, entity := range likes {
		likesDomain = append(likesDomain, entity.ToMarketItemLike())
	}

	return likesDomain, nil
}

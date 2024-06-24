package postgres

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	"github.com/legocy-co/legocy/internal/data/postgres"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/domain/collections/models"
	auth "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

type CollectionPostgresRepository struct {
	conn d.Storage
}

func NewCollectionPostgresRepository(conn d.Storage) CollectionPostgresRepository {
	return CollectionPostgresRepository{conn: conn}
}

func (r CollectionPostgresRepository) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &postgres.ErrConnectionLost
	}

	var userLegoSetsDB []*entities.UserLegoSetPostgres

	res := db.Model(&entities.UserLegoSetPostgres{}).
		Preload("User").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Find(&userLegoSetsDB, "user_id = ?", userID)

	if res.RowsAffected == 0 {
		_error := errors.NewAppError(errors.NotFoundError, "No sets found for user")
		return nil, &_error
	}

	if res.Error != nil {
		_error := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &_error
	}

	legoSetsDomain := make([]models.CollectionLegoSet, 0, len(userLegoSetsDB))
	var user *auth.User = nil

	if len(userLegoSetsDB) == 0 {
		_error := errors.NewAppError(errors.NotFoundError, "No sets found for user")
		return nil, &_error
	}

	for _, legoSetDB := range userLegoSetsDB {
		legoSetsDomain = append(legoSetsDomain, legoSetDB.ToCollectionLegoSet())

		if user == nil {
			user = legoSetDB.User.ToUser()
		}
	}

	return &models.LegoCollection{User: *user, Sets: legoSetsDomain}, nil
}

func (r CollectionPostgresRepository) AddSetToUserCollection(
	c context.Context, userID int, collectionSet models.CollectionLegoSetValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.GetCreatedUserLegoSet(&collectionSet, userID)
	result := tx.Create(entity)
	if result.Error != nil {
		tx.Rollback()
		_error := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &_error
	}

	tx.Commit()

	return nil
}

func (r CollectionPostgresRepository) RemoveSetFromUserCollection(
	c context.Context, userID int, collectionSetID int) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &postgres.ErrConnectionLost
	}

	tx := db.Begin()

	res := tx.Delete(&entities.UserLegoSetPostgres{UserID: userID}, collectionSetID)
	if res.Error != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			res.Error.Error(),
		)
		return &appErr
	}

	tx.Commit()
	return nil
}

func (r CollectionPostgresRepository) UpdateUserCollectionSetByID(
	c context.Context, userID int, setID int, collectionSet models.CollectionLegoSetValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &postgres.ErrConnectionLost
	}

	var entity *entities.UserLegoSetPostgres
	res := db.Model(entities.UserLegoSetPostgres{UserID: userID}).First(&entity, setID)
	if res.Error != nil {
		_error := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return &_error
	}

	tx := db.Begin()
	entity = entities.GetUpdatedUserLegoSet(&collectionSet, entity, userID)
	res = tx.Save(entity)

	if res.Error != nil {
		tx.Rollback()
		_error := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return &_error
	}

	tx.Commit()
	return nil
}

func (r CollectionPostgresRepository) GetCollectionSetOwner(c context.Context, collectionSetID int) (int, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return -1, &postgres.ErrConnectionLost
	}

	var ownerID int
	err := db.Model(
		entities.UserLegoSetPostgres{},
	).Where("id=?", collectionSetID).Select("user_id").First(&ownerID).Error

	var _error *errors.AppError
	if err != nil {
		*_error = errors.NewAppError(errors.NotFoundError, err.Error())
	}
	return ownerID, _error
}

func (r CollectionPostgresRepository) GetUserCollectionSetsAmount(userID int) (int, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return 0, &postgres.ErrConnectionLost
	}

	var total int64

	err := db.Model(
		entities.UserLegoSetPostgres{},
	).Where("user_id=?", userID).Count(&total).Error

	if err != nil {
		_error := errors.NewAppError(errors.NotFoundError, err.Error())
		return 0, &_error
	}

	return int(total), nil
}

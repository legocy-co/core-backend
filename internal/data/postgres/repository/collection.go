package postgres

import (
	"context"
	d "legocy-go/internal/data"
	entities "legocy-go/internal/data/postgres/entity"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/errors"
	auth "legocy-go/internal/domain/users/models"
)

type CollectionPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewCollectionPostgresRepository(conn d.DataBaseConnection) CollectionPostgresRepository {
	return CollectionPostgresRepository{conn: conn}
}

func (r CollectionPostgresRepository) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var userLegoSetsDB []*entities.UserLegoSetPostgres

	res := db.Model(&entities.UserLegoSetPostgres{}).
		Preload("User").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").
		Find(&userLegoSetsDB, "user_id = ?", userID)

	if res.Error != nil {
		_error := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return nil, &_error
	}

	legoSetsDomain := make([]models.CollectionLegoSet, 0, len(userLegoSetsDB))
	var user *auth.User
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
		return &d.ErrConnectionLost
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
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	res := tx.Delete(&entities.UserLegoSetPostgres{UserID: userID}, collectionSetID)
	if res != nil {
		tx.Rollback()
		_error := errors.NewAppError(errors.ConflictError, res.Error.Error())
		return &_error
	}

	tx.Commit()
	return nil
}

func (r CollectionPostgresRepository) UpdateUserCollectionSetByID(
	c context.Context, userID int, setID int, collectionSet models.CollectionLegoSetValueObject) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
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
		return -1, &d.ErrConnectionLost
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

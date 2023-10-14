package postgres

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entity"
	"legocy-go/internal/domain/collections/models"
	auth "legocy-go/internal/domain/users/models"
)

type CollectionPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewCollectionPostgresRepository(conn d.DataBaseConnection) CollectionPostgresRepository {
	return CollectionPostgresRepository{conn: conn}
}

func (r CollectionPostgresRepository) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, error) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	var userLegoSetsDB []*entities.UserLegoSetPostgres

	res := db.Model(&entities.UserLegoSetPostgres{}).
		Preload("User").
		Preload("LegoSet").Preload("LegoSet.LegoSeries").
		Preload("Currency").
		Find(&userLegoSetsDB, "user_id = ?", userID)

	if res.Error != nil {
		return nil, res.Error
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
	c context.Context, userID int, collectionSet models.CollectionLegoSetValueObject) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	entity := entities.GetCreatedUserLegoSet(&collectionSet, userID)
	result := tx.Create(entity)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}

func (r CollectionPostgresRepository) RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}

	tx := db.Begin()

	res := tx.Delete(&entities.UserLegoSetPostgres{UserID: userID}, collectionSetID)
	if res != nil {
		tx.Rollback()
		return res.Error
	}

	tx.Commit()
	return nil
}

func (r CollectionPostgresRepository) UpdateUserCollectionSetByID(c context.Context, userID int, setID int, collectionSet models.CollectionLegoSetValueObject) error {

	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	var entity *entities.UserLegoSetPostgres
	res := db.Model(entities.UserLegoSetPostgres{UserID: userID}).First(&entity, setID)
	if res.Error != nil {
		return res.Error
	}

	tx := db.Begin()
	entity = entities.GetUpdatedUserLegoSet(&collectionSet, entity, userID)
	res = tx.Save(entity)

	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	tx.Commit()
	return nil
}

package repository

import (
	"context"
	"legocy-go/internal/domain/collections/models"
)

type UserCollectionRepository interface {
	GetUserCollection(c context.Context, userID int) (*models.LegoCollection, error)
	AddSetToUserCollection(c context.Context, userID int, collectionSet models.CollectionLegoSetValueObject) error
	RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) error
	UpdateUserCollectionSetByID(c context.Context, userID int, setID int, collectionSet models.CollectionLegoSetValueObject) error
}

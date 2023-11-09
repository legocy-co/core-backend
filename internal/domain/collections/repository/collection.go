package repository

import (
	"context"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/errors"
)

type UserCollectionRepository interface {
	GetUserCollection(c context.Context, userID int) (*models.LegoCollection, *errors.AppError)
	GetCollectionSetOwner(c context.Context, collectionSetID int) (int, *errors.AppError)
	AddSetToUserCollection(c context.Context, userID int, collectionSet models.CollectionLegoSetValueObject) *errors.AppError
	RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) *errors.AppError
	UpdateUserCollectionSetByID(c context.Context, userID int, setID int, collectionSet models.CollectionLegoSetValueObject) *errors.AppError
}

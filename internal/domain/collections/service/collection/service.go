package service

import (
	"context"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/collections/repository"
)

type UserCollectionService struct {
	r repository.UserCollectionRepository
}

func NewUserCollectionService(r repository.UserCollectionRepository) UserCollectionService {
	return UserCollectionService{r: r}
}

func (s UserCollectionService) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, error) {
	return s.r.GetUserCollection(c, userID)
}

func (s UserCollectionService) AddSetToUserCollection(c context.Context, userID int, vo models.CollectionLegoSetValueObject) error {
	return s.r.AddSetToUserCollection(c, userID, vo)
}

func (s UserCollectionService) RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) error {
	return s.r.RemoveSetFromUserCollection(c, userID, collectionSetID)
}

func (s UserCollectionService) UpdateUserCollectionSet(c context.Context, userID int, collectionSetID int, vo models.CollectionLegoSetValueObject) error {
	return s.r.UpdateUserCollectionSetByID(c, userID, collectionSetID, vo)
}

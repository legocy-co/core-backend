package service

import (
	"context"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/collections/repository"
)

type UserCollectionService struct {
	collectionRepository repository.UserCollectionRepository
	valuationRepository  repository.LegoSetValuationRepository
}

func NewUserCollectionService(r repository.UserCollectionRepository) UserCollectionService {
	return UserCollectionService{collectionRepository: r}
}

func (s UserCollectionService) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, error) {
	return s.collectionRepository.GetUserCollection(c, userID)
}

func (s UserCollectionService) AddSetToUserCollection(c context.Context, userID int, vo models.CollectionLegoSetValueObject) error {
	return s.collectionRepository.AddSetToUserCollection(c, userID, vo)
}

func (s UserCollectionService) RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) error {
	return s.collectionRepository.RemoveSetFromUserCollection(c, userID, collectionSetID)
}

func (s UserCollectionService) UpdateUserCollectionSet(c context.Context, userID int, collectionSetID int, vo models.CollectionLegoSetValueObject) error {
	return s.collectionRepository.UpdateUserCollectionSetByID(c, userID, collectionSetID, vo)
}

func (s UserCollectionService) GetUserCollectionValuation(c context.Context, userID int, currencyID int) ([]models.LegoSetValuation, error) {
	userCollection, err := s.GetUserCollection(c, userID)
	if err != nil {
		return []models.LegoSetValuation{}, err
	}

	var setValuations []models.LegoSetValuation

	for _, userLegoSet := range userCollection.Sets {
		setValuation, err := s.valuationRepository.GetLegoSetValuationBySetStateCurrency(
			c, userLegoSet.LegoSet.ID, userLegoSet.CurrentState, currencyID,
		)

		if err != nil {
			return []models.LegoSetValuation{}, err
		}

		setValuations = append(setValuations, *setValuation)
	}

	return setValuations, nil
}

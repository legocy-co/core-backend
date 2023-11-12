package service

import (
	"context"
	"legocy-go/internal/app/errors"
	calculator "legocy-go/internal/domain/calculator/models"
	repository2 "legocy-go/internal/domain/calculator/repository"
	"legocy-go/internal/domain/collections/models"
	"legocy-go/internal/domain/collections/repository"
	auth "legocy-go/internal/domain/users/models"
	users "legocy-go/internal/domain/users/repository"
)

type UserCollectionService struct {
	collectionRepository repository.UserCollectionRepository
	valuationRepository  repository2.LegoSetValuationRepository
	usersRepository      users.UserRepository
}

func NewUserCollectionService(
	collectionRepo repository.UserCollectionRepository,
	valuationRepo repository2.LegoSetValuationRepository,
	usersRepo users.UserRepository) UserCollectionService {
	return UserCollectionService{
		collectionRepository: collectionRepo,
		valuationRepository:  valuationRepo,
		usersRepository:      usersRepo,
	}
}

func (s UserCollectionService) GetUserCollection(c context.Context, userID int) (*models.LegoCollection, *errors.AppError) {
	return s.collectionRepository.GetUserCollection(c, userID)
}

func (s UserCollectionService) AddSetToUserCollection(c context.Context, userID int, vo models.CollectionLegoSetValueObject) *errors.AppError {
	return s.collectionRepository.AddSetToUserCollection(c, userID, vo)
}

func (s UserCollectionService) RemoveSetFromUserCollection(c context.Context, userID int, collectionSetID int) *errors.AppError {
	return s.collectionRepository.RemoveSetFromUserCollection(c, userID, collectionSetID)
}

func (s UserCollectionService) UpdateUserCollectionSet(c context.Context, userID int, collectionSetID int, vo models.CollectionLegoSetValueObject) *errors.AppError {
	return s.collectionRepository.UpdateUserCollectionSetByID(c, userID, collectionSetID, vo)
}

func (s UserCollectionService) GetUserCollectionValuation(c context.Context, userID int, currencyID int) ([]calculator.LegoSetValuation, *auth.User, *errors.AppError) {
	userCollection, err := s.GetUserCollection(c, userID)
	if err != nil {
		return []calculator.LegoSetValuation{}, nil, err
	}

	var setValuations []calculator.LegoSetValuation

	for _, userLegoSet := range userCollection.Sets {
		setValuation, err := s.valuationRepository.GetLegoSetValuationBySetStateCurrency(
			c, userLegoSet.LegoSet.ID, userLegoSet.CurrentState, currencyID,
		)

		if err != nil {
			return []calculator.LegoSetValuation{}, nil, err
		}

		setValuations = append(setValuations, *setValuation)
	}

	user, err := s.usersRepository.GetUserByID(c, userID)
	if err != nil {
		return nil, nil, err
	}

	return setValuations, user, nil
}

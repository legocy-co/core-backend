package service

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	calculator "github.com/legocy-co/legocy/internal/domain/calculator/models"
	c "github.com/legocy-co/legocy/internal/domain/calculator/repository"
	"github.com/legocy-co/legocy/internal/domain/collections/models"
	"github.com/legocy-co/legocy/internal/domain/collections/repository"
	users "github.com/legocy-co/legocy/internal/domain/users/models"
	u "github.com/legocy-co/legocy/internal/domain/users/repository"
)

type UserCollectionService struct {
	collectionRepository repository.UserCollectionRepository
	valuationRepository  c.LegoSetValuationRepository
	usersRepository      u.UserRepository
}

func NewUserCollectionService(
	collectionRepo repository.UserCollectionRepository,
	valuationRepo c.LegoSetValuationRepository,
	usersRepo u.UserRepository) UserCollectionService {
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

func (s UserCollectionService) GetUserCollectionValuation(c context.Context, userID int) ([]calculator.LegoSetValuation, *users.User, *errors.AppError) {
	userCollection, err := s.GetUserCollection(c, userID)
	if err != nil {
		return []calculator.LegoSetValuation{}, nil, err
	}

	var setValuations []calculator.LegoSetValuation

	for _, userLegoSet := range userCollection.Sets {
		setValuation, err := s.valuationRepository.GetLegoSetValuationBySetStateCurrency(
			c, userLegoSet.LegoSet.ID, userLegoSet.CurrentState,
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

package service

import "legocy-go/internal/domain/collections/repository"

type UserCollectionService struct {
	r repository.UserCollectionRepository
}

func NewUserCollectionService(r repository.UserCollectionRepository) UserCollectionService {
	return UserCollectionService{r: r}
}

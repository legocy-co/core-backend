package usecase

import (
	"legocy-go/pkg/auth/repository"
)

type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: r,
	}
}

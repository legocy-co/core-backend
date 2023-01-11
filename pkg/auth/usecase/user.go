package usecase

import (
	r "legocy-go/pkg/auth/repository"
)

type UserUseCase struct {
	userRepo r.UserRepository
}

func NewUserUseCase(repo r.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

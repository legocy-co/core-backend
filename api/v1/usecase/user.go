package v1

import (
	"context"
	res "legocy-go/api/v1/resources"
	models "legocy-go/pkg/auth/models"
	r "legocy-go/pkg/auth/repository"
)

type UserUseCase struct {
	repo r.UserRepository
}

func NewUserUsecase(repo r.UserRepository) UserUseCase {
	return UserUseCase{repo: repo}
}

func (s *UserUseCase) ValidateUser(c context.Context, req res.JWTRequest) error {
	return s.repo.ValidateUser(c, req.Email, req.Password)
}

func (s *UserUseCase) CreateUser(c context.Context, u *models.User, password string) error {
	return s.repo.CreateUser(c, u, password)
}

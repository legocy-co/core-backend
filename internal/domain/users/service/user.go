package service

import (
	"context"
	res "legocy-go/internal/delievery/http/resources/users"
	models "legocy-go/internal/domain/users/models"
	r "legocy-go/internal/domain/users/repository"
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

func (s *UserUseCase) GetUserByEmail(c context.Context, email string) (*models.User, error) {
	return s.repo.GetUserByEmail(c, email)
}

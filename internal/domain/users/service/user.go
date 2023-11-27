package service

import (
	"context"
	"legocy-go/internal/app/errors"
	res "legocy-go/internal/delivery/http/resources/users"
	models "legocy-go/internal/domain/users/models"
	r "legocy-go/internal/domain/users/repository"
)

type UserService struct {
	repo r.UserRepository
}

func NewUserService(repo r.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) ValidateUser(c context.Context, req res.JWTRequest) *errors.AppError {
	return s.repo.ValidateUser(c, req.Email, req.Password)
}

func (s *UserService) CreateUser(c context.Context, u *models.User, password string) *errors.AppError {
	return s.repo.CreateUser(c, u, password)
}

func (s *UserService) GetUserByEmail(c context.Context, email string) (*models.User, *errors.AppError) {
	return s.repo.GetUserByEmail(c, email)
}

func (s *UserService) GetUserByID(c context.Context, id int) (*models.User, *errors.AppError) {
	return s.repo.GetUserByID(c, id)
}

package service

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	res "github.com/legocy-co/legocy/internal/delivery/http/schemas/users"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	r "github.com/legocy-co/legocy/internal/domain/users/repository"
)

type UserService struct {
	repo r.UserRepository
}

func NewUserService(repo r.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) ValidateUserCredentials(c context.Context, req res.SignInRequest) *errors.AppError {
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

func (s *UserService) UpdateUser(id int, vo models.UserValueObject) *errors.AppError {
	return s.repo.UpdateUser(id, vo)
}

package admin

import (
	"context"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	r "github.com/legocy-co/legocy/internal/domain/users/repository"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

type UserAdminService struct {
	repo r.UserAdminRepository
}

func NewUserAdminService(repo r.UserAdminRepository) UserAdminService {
	return UserAdminService{repo: repo}
}

func (s UserAdminService) GetUsers(c context.Context) ([]*models.UserAdmin, *errors.AppError) {
	return s.repo.GetUsers(c)
}

func (s UserAdminService) GetUserByID(c context.Context, id int) (*models.UserAdmin, *errors.AppError) {
	return s.repo.GetUserByID(c, id)
}

func (s UserAdminService) CreateAdmin(
	c context.Context, ua *models.UserAdminValueObject, password string) *errors.AppError {
	return s.repo.CreateAdmin(c, ua, password)
}

func (s UserAdminService) DeleteUser(c context.Context, id int) *errors.AppError {
	return s.repo.DeleteUser(c, id)
}

func (s UserAdminService) UpdateUser(c context.Context, id int, vo *models.UserAdminValueObject) (*models.UserAdmin, *errors.AppError) {
	return s.repo.UpdateUserByID(c, id, vo)
}

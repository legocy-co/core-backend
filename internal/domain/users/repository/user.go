package repository

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserRepository interface {
	CreateUser(c context.Context, u *models.User, password string) *errors.AppError
	ValidateUser(c context.Context, email, password string) *errors.AppError
	GetUsers(c context.Context) ([]*models.User, *errors.AppError)
	GetUserByEmail(c context.Context, email string) (*models.User, *errors.AppError)
	GetUserByID(c context.Context, id int) (*models.User, *errors.AppError)
	DeleteUser(c context.Context, id int) *errors.AppError
}

type UserAdminRepository interface {
	CreateAdmin(c context.Context, ua *models.UserAdmin, password string) *errors.AppError
	GetUsers(c context.Context) ([]*models.UserAdmin, *errors.AppError)
	GetUserByID(c context.Context, id int) (*models.UserAdmin, *errors.AppError)
	DeleteUser(c context.Context, id int) *errors.AppError
	UpdateUserByID(c context.Context, itemId int, vo *models.UserAdminValueObject) (*models.UserAdmin, *errors.AppError)
}

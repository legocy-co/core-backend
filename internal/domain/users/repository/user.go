package repository

import (
	"context"
	models "legocy-go/internal/domain/users/models"
	"legocy-go/internal/domain/users/models/admin"
)

type UserRepository interface {
	CreateUser(c context.Context, u *models.User, password string) error
	ValidateUser(c context.Context, email, password string) error
	GetUsers(c context.Context) ([]*models.User, error)
	GetUserByEmail(c context.Context, email string) (*models.User, error)
	GetUserByID(c context.Context, id int) (*models.User, error)
	DeleteUser(c context.Context, id int) error
}

type UserAdminRepository interface {
	CreateAdmin(c context.Context, u *admin.UserAdmin, password string) error
	GetUserByEmail(c context.Context, email string) (*admin.UserAdmin, error)
	GetUserByID(c context.Context, id int) (*admin.UserAdmin, error)
	DeleteUser(c context.Context, id int) error
}

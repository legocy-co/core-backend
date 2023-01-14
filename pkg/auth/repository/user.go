package repository

import (
	"context"
	models "legocy-go/pkg/auth/models"
)

type UserRepository interface {
	CreateUser(c context.Context, u *models.User, password string) error
	ValidateUser(c context.Context, email, password string) error
	GetUsers(c context.Context) ([]*models.User, error)
	GetUser(c context.Context, id int) (*models.User, error)
	DeleteUser(c context.Context, id int) error
}

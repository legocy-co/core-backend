package repository

import (
	"context"
	models "legocy-go/pkg/auth/models"
)

type UserRepository interface {
	CreateUser(c *context.Context) error
	GetUser(c *context.Context) (*models.User, error)
	DeleteUser(c *context.Context, id int) error
}

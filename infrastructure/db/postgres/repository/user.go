package postgres

import (
	"context"
	h "legocy-go/helpers"
	p "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	models "legocy-go/pkg/auth/models"
)

type UserPostgresRepository struct {
	conn *p.PostrgresConnection
}

func (r *UserPostgresRepository) CreateUser(c context.Context, u *models.User, password string) error {
	db := r.conn.GetDB()

	if db == nil {
		return p.ErrConnectionLost
	}
	passwordHash, err := h.HashPassword(password)
	if err != nil {
		return h.ErrHashError
	}

	var entity entities.UserPostgres = *entities.FromUser(u, passwordHash)

	db.Create(&entity)
	return nil
}

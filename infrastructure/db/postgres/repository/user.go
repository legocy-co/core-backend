package postgres

import (
	"context"
	h "legocy-go/helpers"
	d "legocy-go/infrastructure/db"
	p "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	err "legocy-go/pkg/auth/errors"
	models "legocy-go/pkg/auth/models"
)

type UserPostgresRepository struct {
	conn *p.PostrgresConnection
}

func (r *UserPostgresRepository) CreateUser(c context.Context, u *models.User, password string) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}
	passwordHash, err := h.HashPassword(password)
	if err != nil {
		return h.ErrHashError
	}

	var entity entities.UserPostgres = *entities.FromUser(u, passwordHash)

	db.Create(&entity)
	return nil
}

func (r *UserPostgresRepository) ValidateUser(c context.Context, username, email, password string) error {

	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	db.Model(entities.UserPostgres{Username: username, Email: email}).First(entity)

	if entity == nil {
		return err.ErrUserNotFound
	}

	passwordsMacthed := h.CheckPasswordHash(password, entity.Password)
	if !passwordsMacthed {
		return err.ErrWrongPassword
	}

	return nil
}

func (r *UserPostgresRepository) GetUsers(c context.Context) ([]*models.User, error) {
	var usersDb []*entities.UserPostgres
	var users []*models.User

	db := r.conn.GetDB()
	if db == nil {
		return users, d.ErrConnectionLost
	}

	db.Find(usersDb)

	for _, userDb := range usersDb {
		users = append(users, userDb.ToUser())
	}

	var errOutput error

	if len(users) == 0 {
		errOutput = err.ErrUserNotFound
	}

	return users, errOutput
}

func (r *UserPostgresRepository) GetUser(c context.Context, id int) (*models.User, error) {
	var user *models.User
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, d.ErrConnectionLost
	}

	db.First(entity, id)
	if entity == nil {
		return user, err.ErrUserNotFound
	}

	user = entity.ToUser()
	return user, nil
}

func (r *UserPostgresRepository) DeleteUser(c context.Context, id int) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.UserPostgres{}, id)
	return nil

}

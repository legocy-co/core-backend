package postgres

import (
	"context"
	"fmt"
	h "legocy-go/helpers"
	d "legocy-go/infrastructure/db"
	p "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	e "legocy-go/pkg/auth/errors"
	models "legocy-go/pkg/auth/models"
)

type UserPostgresRepository struct {
	conn *p.PostrgresConnection
}

func NewUserPostgresRepository(conn *p.PostrgresConnection) UserPostgresRepository {
	return UserPostgresRepository{conn: conn}
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

	result := db.Create(&entity)
	if result.Error != nil {
		return e.ErrUserAlreadyExists
	}
	db.Commit()

	return nil
}

func (r *UserPostgresRepository) ValidateUser(c context.Context, email, password string) error {

	db := r.conn.GetDB()
	if db == nil {
		fmt.Println("Error Connecting to database!")
		return d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	db.Model(entities.UserPostgres{}).First(&entity, entities.UserPostgres{Email: email})

	if entity == nil {
		fmt.Printf("User with email = %v not found", email)
		return e.ErrUserNotFound
	}

	passwordsMacthed := h.CheckPasswordHash(password, entity.Password)
	if !passwordsMacthed {
		return e.ErrWrongPassword
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
		errOutput = e.ErrUserNotFound
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
		return user, e.ErrUserNotFound
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
	db.Commit()
	return nil

}

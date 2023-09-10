package postgres

import (
	"context"
	"fmt"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entity"
	e "legocy-go/internal/domain/users/errors"
	models "legocy-go/internal/domain/users/models"
	h "legocy-go/pkg/helpers"
)

type UserPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserPostgresRepository(conn d.DataBaseConnection) UserPostgresRepository {
	return UserPostgresRepository{conn: conn}
}

func (r UserPostgresRepository) CreateUser(c context.Context, u *models.User, password string) error {
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
	return result.Error
}

func (r UserPostgresRepository) ValidateUser(c context.Context, email, password string) error {

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

func (r UserPostgresRepository) GetUsers(c context.Context) ([]*models.User, error) {
	var usersDb []*entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, d.ErrConnectionLost
	}

	db.Find(usersDb)

	users := make([]*models.User, 0, len(usersDb))
	for _, userDb := range usersDb {
		users = append(users, userDb.ToUser())
	}

	var errOutput error

	if len(users) == 0 {
		errOutput = e.ErrUserNotFound
	}

	return users, errOutput
}

func (r UserPostgresRepository) GetUserByEmail(c context.Context, email string) (*models.User, error) {
	var user *models.User
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, d.ErrConnectionLost
	}

	db.Where("email = ?", email).First(&entity)
	if entity == nil {
		return user, e.ErrUserNotFound
	}

	user = entity.ToUser()
	return user, nil
}

func (r UserPostgresRepository) GetUserByID(c context.Context, id int) (*models.User, error) {
	var user *models.User
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, d.ErrConnectionLost
	}

	db.First(&entity, id)
	if entity == nil {
		return user, e.ErrUserNotFound
	}

	user = entity.ToUser()
	return user, nil
}

func (r UserPostgresRepository) DeleteUser(c context.Context, id int) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.UserPostgres{}, id)
	db.Commit()
	return nil

}

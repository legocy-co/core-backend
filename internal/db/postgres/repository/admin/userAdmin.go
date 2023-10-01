package admin

import (
	"context"
	d "legocy-go/internal/db"
	entities "legocy-go/internal/db/postgres/entity"
	e "legocy-go/internal/domain/users/errors"
	"legocy-go/internal/domain/users/models/admin"
	h "legocy-go/pkg/helpers"
)

type UserAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserAdminPostgresRepository(conn d.DataBaseConnection) UserAdminPostgresRepository {
	return UserAdminPostgresRepository{conn: conn}
}

func (r UserAdminPostgresRepository) CreateAdmin(
	c context.Context, ua *admin.UserAdmin, password string) error {
	db := r.conn.GetDB()

	if db == nil {
		return d.ErrConnectionLost
	}
	passwordHash, err := h.HashPassword(password)
	if err != nil {
		return h.ErrHashError
	}

	var entity entities.UserPostgres = *entities.FromAdmin(ua, passwordHash)

	result := db.Create(&entity)
	return result.Error
}

func (r UserAdminPostgresRepository) GetUserByEmail(
	c context.Context, email string) (*admin.UserAdmin, error) {
	var user *admin.UserAdmin
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, d.ErrConnectionLost
	}

	db.Where("email = ?", email).First(&entity)
	if entity == nil {
		return user, e.ErrUserNotFound
	}

	user = entity.ToUserAdmin()
	return user, nil
}

func (r UserAdminPostgresRepository) GetUserByID(c context.Context, id int) (*admin.UserAdmin, error) {
	var user *admin.UserAdmin
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, d.ErrConnectionLost
	}

	db.First(&entity, id)
	if entity == nil {
		return user, e.ErrUserNotFound
	}

	user = entity.ToUserAdmin()
	return user, nil
}

func (r UserAdminPostgresRepository) DeleteUser(c context.Context, id int) error {
	db := r.conn.GetDB()
	if db == nil {
		return d.ErrConnectionLost
	}

	db.Delete(&entities.UserPostgres{}, id)
	db.Commit()
	return nil
}

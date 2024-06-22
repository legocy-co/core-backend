package google

import (
	"context"
	postgres2 "github.com/legocy-co/legocy/internal/data/postgres"
	"github.com/legocy-co/legocy/internal/pkg/errors"

	d "github.com/legocy-co/legocy/internal/data"
	entity "github.com/legocy-co/legocy/internal/data/postgres/entity"
	postgres "github.com/legocy-co/legocy/internal/data/postgres/repository"
	e "github.com/legocy-co/legocy/internal/domain/users/errors"
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/helpers"
)

type UserAuthRepository struct {
	conn d.Storage
	base postgres.UserPostgresRepository
}

func NewUserAuthRepository(conn d.Storage, dispatcher events.Dispatcher) UserAuthRepository {
	return UserAuthRepository{
		conn: conn,
		base: postgres.NewUserPostgresRepository(conn, dispatcher),
	}
}

func (r UserAuthRepository) GetByExternalID(c context.Context, externalID string) (*models.User, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &postgres2.ErrConnectionLost
	}

	var userDB *entity.UserPostgres
	if ok := db.Where("google_id = ?", externalID).First(&userDB).RowsAffected > 0; !ok {
		return nil, &e.ErrUserNotFound
	}

	return userDB.ToUser(), nil
}

func (r UserAuthRepository) CreateUser(c context.Context, u models.UserValueObject) *errors.AppError {

	user := models.FromVO(u)
	password := helpers.GetRandomPassword()

	return r.base.CreateUser(c, user, password)
}

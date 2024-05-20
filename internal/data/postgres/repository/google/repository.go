package google

import (
	"context"

	d "github.com/legocy-co/legocy/internal/data"
	entity "github.com/legocy-co/legocy/internal/data/postgres/entity"
	postgres "github.com/legocy-co/legocy/internal/data/postgres/repository"
	e "github.com/legocy-co/legocy/internal/domain/users/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/internal/pkg/events"
	"github.com/legocy-co/legocy/pkg/helpers"
)

type UserAuthRepository struct {
	conn d.DBConn
	base postgres.UserPostgresRepository
}

func NewUserAuthRepository(conn d.DBConn, dispatcher events.Dispatcher) UserAuthRepository {
	return UserAuthRepository{
		conn: conn,
		base: postgres.NewUserPostgresRepository(conn, dispatcher),
	}
}

func (r UserAuthRepository) GetByExternalID(c context.Context, externalID string) (*models.User, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
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

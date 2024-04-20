package postgres

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/delivery/kafka/types/users"
	e "github.com/legocy-co/legocy/internal/domain/users/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	h "github.com/legocy-co/legocy/pkg/helpers"
	"github.com/legocy-co/legocy/pkg/kafka"
)

type UserPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserPostgresRepository(conn d.DataBaseConnection) UserPostgresRepository {
	return UserPostgresRepository{conn: conn}
}

func (r UserPostgresRepository) CreateUser(c context.Context, u *models.User, password string) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}
	passwordHash, err := h.HashPassword(password)
	if err != nil {
		return &h.ErrHashError
	}

	tx := db.Begin()

	entity := *entities.FromUser(u, passwordHash)

	if err := tx.Create(entity).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	u = entity.ToUser()

	if err := kafka.ProduceJSONEvent(kafka.UserCreatedTopic, users.FromDomain(u)); err != nil {
		appErr := errors.NewAppError(errors.InternalError, err.Error())
		return &appErr
	}

	return nil
}

func (r UserPostgresRepository) UpdateUser(id int, vo models.UserValueObject) *errors.AppError {
	db := r.conn.GetDB()

	if db == nil {
		return &d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	db.First(&entity, id)

	if entity == nil {
		return &e.ErrUserNotFound
	}

	entity = entities.GetUpdatedUserEntity(vo, entity)
	err := db.Save(&entity).Error

	if err != nil {
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	eventData := users.FromDomainVO(&vo, id)
	if err := kafka.ProduceJSONEvent(kafka.UserUpdatedTopic, eventData); err != nil {
		appErr := errors.NewAppError(errors.InternalError, err.Error())
		return &appErr
	}

	return nil
}

func (r UserPostgresRepository) ValidateUser(c context.Context, email, password string) *errors.AppError {

	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	db.Model(entities.UserPostgres{}).First(&entity, entities.UserPostgres{Email: email})

	if entity == nil {
		return &e.ErrUserNotFound
	}

	passwordsMatched := h.CheckPasswordHash(password, entity.Password)
	if !passwordsMatched {
		return &e.ErrWrongPassword
	}

	return nil
}

func (r UserPostgresRepository) GetUsers(c context.Context) ([]*models.User, *errors.AppError) {
	var usersDb []*entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	db.Preload("Images").Find(usersDb)

	users := make([]*models.User, 0, len(usersDb))
	for _, userDb := range usersDb {
		users = append(users, userDb.ToUser())
	}

	var errOutput *errors.AppError

	if len(users) == 0 {
		errOutput = &e.ErrUserNotFound
	}

	return users, errOutput
}

func (r UserPostgresRepository) GetUserByEmail(c context.Context, email string) (*models.User, *errors.AppError) {
	var user *models.User
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, &d.ErrConnectionLost
	}

	db.Where("email = ?", email).First(&entity)
	if entity == nil {
		return user, &e.ErrUserNotFound
	}

	user = entity.ToUser()
	return user, nil
}

func (r UserPostgresRepository) GetUserByID(c context.Context, id int) (*models.User, *errors.AppError) {
	var user *models.User
	var entity *entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return user, &d.ErrConnectionLost
	}

	db.Preload("Images").First(&entity, id)
	if entity == nil {
		return user, &e.ErrUserNotFound
	}

	user = entity.ToUser()
	return user, nil
}

func (r UserPostgresRepository) DeleteUser(c context.Context, id int) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	err := tx.Delete(&entities.UserPostgres{}, id).Error
	if err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, err.Error())
		return &appErr
	}

	tx.Commit()
	return nil
}

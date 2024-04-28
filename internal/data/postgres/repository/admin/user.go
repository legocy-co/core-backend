package admin

import (
	"context"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	"github.com/legocy-co/legocy/internal/delivery/kafka/types/users"
	e "github.com/legocy-co/legocy/internal/domain/users/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/app/errors"
	"github.com/legocy-co/legocy/internal/pkg/events"
	h "github.com/legocy-co/legocy/pkg/helpers"
	"github.com/legocy-co/legocy/pkg/kafka"
)

type UserAdminPostgresRepository struct {
	conn       d.DBConn
	dispatcher events.Dispatcher
}

func NewUserAdminPostgresRepository(conn d.DBConn, dispatcher events.Dispatcher) UserAdminPostgresRepository {
	return UserAdminPostgresRepository{
		conn:       conn,
		dispatcher: dispatcher,
	}
}

func (r UserAdminPostgresRepository) GetUsers(c context.Context) ([]*models.UserAdmin, *errors.AppError) {
	var usersAdminDb []*entities.UserPostgres

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	db.Find(&usersAdminDb)

	users := make([]*models.UserAdmin, 0, len(usersAdminDb))
	for _, usersAdminDb := range usersAdminDb {
		users = append(users, usersAdminDb.ToUserAdmin())
	}

	var errOutput *errors.AppError

	if len(users) == 0 {
		errOutput = &e.ErrUserNotFound
	}

	return users, errOutput
}

func (r UserAdminPostgresRepository) GetUserByID(
	c context.Context, id int) (*models.UserAdmin, *errors.AppError) {
	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var userAdmin *models.UserAdmin

	var entity *entities.UserPostgres
	ok := db.First(&entity, id).RowsAffected > 0
	if !ok {
		return userAdmin, &e.ErrUserNotFound
	}

	userAdmin = entity.ToUserAdmin()
	return userAdmin, nil
}

func (r UserAdminPostgresRepository) GetUserByEmail(
	c context.Context, email string) (*models.UserAdmin, *errors.AppError) {

	db := r.conn.GetDB()
	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	ok := db.Where("email = ?", email).First(&entity).RowsAffected > 0
	if !ok {
		return nil, &e.ErrUserNotFound
	}

	return entity.ToUserAdmin(), nil
}

func (r UserAdminPostgresRepository) CreateAdmin(c context.Context, ua *models.UserAdminValueObject, password string) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	passwordHash, appErr := h.HashPassword(password)
	if appErr != nil {
		return appErr
	}

	var entity = *entities.FromAdminVO(ua, passwordHash)
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			err.Error(),
		)
		return &appErr
	}

	tx.Commit()

	eventData := users.FromDomainVOAdmin(ua, int(entity.ID))
	if err := r.dispatcher.ProduceJSONEvent(kafka.UserCreatedTopic, eventData); err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.InternalError,
			err.Error(),
		)
		return &appErr
	}

	return nil
}

func (r UserAdminPostgresRepository) UpdateUserByID(
	c context.Context, userId int, vo *models.UserAdminValueObject) (*models.UserAdmin, *errors.AppError) {
	db := r.conn.GetDB()

	if db == nil {
		return nil, &d.ErrConnectionLost
	}

	var entity *entities.UserPostgres
	_ = db.First(&entity, userId)
	if entity == nil {
		return nil, &e.ErrUserNotFound
	}

	tx := db.Begin()

	entityUpdated := entity.GetUpdatedUserAdmin(*vo)
	if err := tx.Save(entityUpdated).Error; err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.ConflictError,
			err.Error(),
		)
		return nil, &appErr
	}

	eventData := users.FromDomainVOAdmin(vo, userId)
	if err := r.dispatcher.ProduceJSONEvent(kafka.UserUpdatedTopic, eventData); err != nil {
		tx.Rollback()
		appErr := errors.NewAppError(
			errors.InternalError,
			err.Error(),
		)
		return nil, &appErr
	}

	tx.Commit()

	return r.GetUserByID(c, userId)
}

func (r UserAdminPostgresRepository) DeleteUser(c context.Context, userId int) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	result := db.Delete(entities.UserPostgres{}, userId)
	if result.Error != nil {
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	return nil
}

func (r UserAdminPostgresRepository) ValidateUser(c context.Context, email, password string) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	var user *entities.UserPostgres

	ok := db.Where("email = ? AND role = ?", email, models.ADMIN).First(&user).RowsAffected > 0
	if !ok {
		return &e.ErrUserNotFound
	}

	passwordsMatch := h.CheckPasswordHash(password, user.Password)
	if !passwordsMatch {
		return &e.ErrInvalidPassword
	}

	return nil
}

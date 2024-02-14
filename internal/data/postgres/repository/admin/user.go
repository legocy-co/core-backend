package admin

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	d "github.com/legocy-co/legocy/internal/data"
	entities "github.com/legocy-co/legocy/internal/data/postgres/entity"
	e "github.com/legocy-co/legocy/internal/domain/users/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	h "github.com/legocy-co/legocy/pkg/helpers"
	"github.com/legocy-co/legocy/pkg/kafka"
)

type UserAdminPostgresRepository struct {
	conn d.DataBaseConnection
}

func NewUserAdminPostgresRepository(conn d.DataBaseConnection) UserAdminPostgresRepository {
	return UserAdminPostgresRepository{conn: conn}
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

	var userAdmin *models.UserAdmin

	var entity *entities.UserPostgres
	ok := db.Where("email = ?", email).First(&entity).RowsAffected > 0
	if !ok {
		return userAdmin, &e.ErrUserNotFound
	}

	userAdmin = entity.ToUserAdmin()
	return userAdmin, nil
}

func (r UserAdminPostgresRepository) CreateAdmin(c context.Context, ua *models.UserAdmin, password string) *errors.AppError {
	db := r.conn.GetDB()
	if db == nil {
		return &d.ErrConnectionLost
	}

	tx := db.Begin()

	passwordHash, err := h.HashPassword(password)
	if err != nil {
		return &h.ErrHashError
	}

	var entity = *entities.FromAdmin(ua, passwordHash)
	result := db.Create(&entity)

	if result.Error != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.ConflictError, result.Error.Error())
		return &appErr
	}

	tx.Commit()

	kafkaErr := kafka.ProduceJSONEvent(kafka.USER_UPDATES_TOPIC, map[string]interface{}{
		"userID": int(entity.ID),
	})
	if kafkaErr != nil {
		tx.Rollback()
		appErr := errors.NewAppError(errors.InternalError, kafkaErr.Error())
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

	entityUpdated := entity.GetUpdatedUserAdmin(*vo)
	db.Save(entityUpdated)

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

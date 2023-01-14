package v1

import (
	res "legocy-go/api/v1/resources"
	psql "legocy-go/infrastructure/db/postgres"
	entities "legocy-go/infrastructure/db/postgres/entities"
	err "legocy-go/pkg/auth/errors"
)

type JWTUseCase struct {
	conn psql.PostrgresConnection
}

func (jwt *JWTUseCase) ValidateUser(u *res.JWTRequest) error {

	db := jwt.conn.GetDB()
	if db == nil {
		return psql.ErrConnectionLost
	}

	record := db.First(&entities.UserPostgres{}, "username = ?", u.Username)
	if record.Error != nil {
		return err.ErrUserNotFound
	}

	return nil
}

package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"legocy-go/internal/app/errors"
)

var ErrHashError = errors.NewAppError(errors.InternalError, "could not hash string")

func HashPassword(password string) (string, *errors.AppError) {
	var err *errors.AppError
	bytes, _e := bcrypt.GenerateFromPassword([]byte(password), 14)

	if _e != nil {
		*err = ErrHashError
	}
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

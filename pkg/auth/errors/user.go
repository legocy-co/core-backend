package auth

import (
	"errors"
)

var ErrUserAlreadyExists = errors.New("user with this credentials already exists")
var ErrUserNotFound = errors.New("user not found")
var ErrWrongPassword = errors.New("wrong password")

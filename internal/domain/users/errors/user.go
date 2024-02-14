package errors

import (
	"github.com/legocy-co/legocy/internal/app/errors"
)

var ErrUserNotFound = errors.NewAppError(errors.NotFoundError, "user not found")
var ErrWrongPassword = errors.NewAppError(errors.PermissionError, "wrong password")
var ErrInvalidImageFilepath = errors.NewAppError(errors.ValidationError, "invalid image filepath")
var ErrInvalidPassword = errors.NewAppError(errors.PermissionError, "invalid password")

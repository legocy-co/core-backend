package errors

import (
	"legocy-go/internal/app/errors"
)

var ErrUserNotFound = errors.NewAppError(errors.NotFoundError, "user not found")
var ErrWrongPassword = errors.NewAppError(errors.PermissionError, "wrong password")
var ErrInvalidImageFilepath = errors.NewAppError(errors.ValidationError, "invalid image filepath")

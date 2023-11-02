package data

import (
	"legocy-go/internal/domain/errors"
)

var ErrConnectionLost = errors.NewAppError(errors.ConflictError, "connection lost")
var ErrConnectionAlreadyExists = errors.NewAppError(errors.ConflictError, "connection already established")
var ErrItemNotFound = errors.NewAppError(errors.NotFoundError, "item(s) not found")

package postgres

import (
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

var ErrConnectionLost = errors.NewAppError(errors.InternalError, "connection lost")
var ErrConnectionAlreadyExists = errors.NewAppError(errors.ConflictError, "connection already established")
var ErrItemNotFound = errors.NewAppError(errors.NotFoundError, "item(s) not found")

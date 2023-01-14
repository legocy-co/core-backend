package postgres

import (
	"errors"
)

var ErrConnectionLost = errors.New("connection lost")
var ErrConnectionAlreadyExists = errors.New("connection already established")

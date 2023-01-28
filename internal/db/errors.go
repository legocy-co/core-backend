package db

import (
	"errors"
)

var ErrConnectionLost = errors.New("connection lost")
var ErrConnectionAlreadyExists = errors.New("connection already established")

var ErrItemNotFound = errors.New("item(s) not found")

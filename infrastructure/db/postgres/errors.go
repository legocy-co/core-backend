package postgres

import (
	"errors"
)

var ErrConnectionLost = errors.New("connection lost")

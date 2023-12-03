package lego

import (
	"github.com/legocy-co/legocy/internal/app/errors"
)

var (
	ErrLegoSetsNotFound   = errors.NewAppError(errors.NotFoundError, "LEGO Set(s) Not Found")
	ErrLegoSeriesNotFound = errors.NewAppError(errors.NotFoundError, "LEGO Series(s) Not Found")
	ErrInvalidLegoState   = errors.NewAppError(errors.ValidationError, "Invalid lego set state value")
)

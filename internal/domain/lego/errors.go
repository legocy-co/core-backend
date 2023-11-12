package lego

import (
	"legocy-go/internal/app/errors"
)

var (
	ErrLegoSetsNotFound   = errors.NewAppError(errors.NotFoundError, "LEGO Set(s) Not Found")
	ErrLegoSeriesNotFound = errors.NewAppError(errors.NotFoundError, "LEGO Series(s) Not Found")
)

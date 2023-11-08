package lego

import "legocy-go/internal/domain/errors"

var (
	ErrLegoSetsNotFound = errors.NewAppError(errors.NotFoundError, "LEGO Set(s) Not Found")
)

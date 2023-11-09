package collections

import (
	"legocy-go/internal/domain/errors"
)

var (
	ErrValuationNotFound = errors.NewAppError(errors.NotFoundError, "LegoSet Valuation Not Found")
)

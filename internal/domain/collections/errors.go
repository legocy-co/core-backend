package collections

import (
	"legocy-go/internal/app/errors"
)

var (
	ErrValuationNotFound = errors.NewAppError(errors.NotFoundError, "LegoSet Valuation Not Found")
)

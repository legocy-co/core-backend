package calculator

import (
	"legocy-go/internal/app/errors"
)

var (
	ErrLegoSetValuationNotFound = errors.NewAppError(errors.NotFoundError, "Lego Set Valuation Not Found")
)

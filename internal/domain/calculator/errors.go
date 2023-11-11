package calculator

import "legocy-go/internal/domain/errors"

var (
	ErrLegoSetValuationNotFound = errors.NewAppError(errors.NotFoundError, "Lego Set Valuation Not Found")
)

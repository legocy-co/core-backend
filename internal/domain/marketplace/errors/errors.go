package errors

import (
	"github.com/legocy-co/legocy/internal/app/errors"
)

var (
	ErrMarketItemsNotFound       = errors.NewAppError(errors.NotFoundError, "marketItems not found")
	ErrUserReviewsNotFound       = errors.NewAppError(errors.NotFoundError, "userReviews not found")
	ErrMarketItemInvalidSellerID = errors.NewAppError(errors.PermissionError, "invalid SellerID")
	ErrMarketItemInvalidStatus   = errors.NewAppError(errors.ValidationError, "invalid status value")
	ErrMarketItemInvalidSetState = errors.NewAppError(errors.ValidationError, "invalid set state")
)

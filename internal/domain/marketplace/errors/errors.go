package errors

import "errors"

var (
	ErrMarketItemsNotFound       = errors.New("marketItems not found")
	ErrUserReviewsNotFound       = errors.New("userReviews not found")
	ErrUsersNotFound             = errors.New("users not found")
	ErrMarketItemInvalidSellerID = errors.New("invalid SellerID")
	ErrMarketItemInvalidStatus   = errors.New("invalid status value")
	ErrMarketItemInvalidSetState = errors.New("invalid set state")
)

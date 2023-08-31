package errors

import "errors"

var (
	ErrMarketItemsNotFound       = errors.New("marketItems not found")
	ErrUserReviewsNotFound       = errors.New("userReviews not found")
	ErrMarketItemInvalidSellerID = errors.New("invalid SellerID")
)

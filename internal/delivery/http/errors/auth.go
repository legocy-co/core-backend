package errors

import (
	"net/http"
)

var ErrTokenHeaderNotFound = NewHttpResponseError(http.StatusBadRequest, "token header not found")
var ErrParsingClaims = NewHttpResponseError(http.StatusBadRequest, "error parsing token claims")

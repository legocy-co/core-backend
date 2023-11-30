package errors

import (
	"net/http"
)

var ErrInvaldTokenHeader = NewHttpResponseError(http.StatusBadRequest, "invalid token header value")
var ErrParsingClaims = NewHttpResponseError(http.StatusBadRequest, "error parsing token claims")

package errors

import (
	"net/http"
)

var ErrParamNotFound = NewHttpResponseError(http.StatusBadRequest, "requested URL param not found")

package errors

import "errors"

var ErrSeriesNotFound = errors.New("LEGO Series Not Found")
var ErrSeriesAlreadyExists = errors.New("LEGO Series with given details already exists")

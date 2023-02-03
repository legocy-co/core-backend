package errors

import "errors"

var ErrTokenHeaderNotFound = errors.New("token header not found")
var ErrParsingClaims = errors.New("error parsing token claims")

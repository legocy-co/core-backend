package errors

import "errors"

var ErrSetNotFound = errors.New("LEGO Set Not Found")
var ErrSetAlreadyExists = errors.New("LEGO Set with given details already exists")

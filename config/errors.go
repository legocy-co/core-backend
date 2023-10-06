package config

import (
	"errors"
)

var ErrConfigAlreadyExists = errors.New("config already exists")
var ErrConfigFileDoesNotExist = errors.New("config file does not exist")

package config

import (
	"errors"
)

var (
	ErrConfigAlreadyExists = errors.New("config already exists")
	ErrConfigNotFound      = errors.New("config not found")
)

package auth

import (
	_ "github.com/legocy-co/legocy/docs"
	ser "github.com/legocy-co/legocy/internal/domain/users/service"
)

type TokenHandler struct {
	service ser.UserService
}

func NewTokenHandler(service ser.UserService) TokenHandler {
	return TokenHandler{service: service}
}

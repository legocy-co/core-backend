package auth

import (
	_ "legocy-go/docs"
	ser "legocy-go/internal/domain/users/service"
)

type TokenHandler struct {
	service ser.UserService
}

func NewTokenHandler(service ser.UserService) TokenHandler {
	return TokenHandler{service: service}
}

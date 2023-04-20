package auth

import (
	_ "legocy-go/docs"
	ser "legocy-go/internal/delievery/http/service/users"
)

type TokenHandler struct {
	service ser.UserUseCase
}

func NewTokenHandler(service ser.UserUseCase) TokenHandler {
	return TokenHandler{service: service}
}

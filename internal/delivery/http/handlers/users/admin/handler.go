package admin

import (
	service "github.com/legocy-co/legocy/internal/domain/users/service/admin"
)

type UserAdminHandler struct {
	service service.UserAdminService
}

func NewUserAdminHandler(service service.UserAdminService) UserAdminHandler {
	return UserAdminHandler{service: service}
}

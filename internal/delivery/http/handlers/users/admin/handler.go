package admin

import (
	service "legocy-go/internal/domain/users/service/admin"
)

type UserAdminHandler struct {
	service service.UserAdminService
}

func NewUserAdminHandler(service service.UserAdminService) UserAdminHandler {
	return UserAdminHandler{service: service}
}

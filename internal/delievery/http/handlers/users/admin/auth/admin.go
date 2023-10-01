package auth

import service "legocy-go/internal/domain/users/service/admin"

type AdminHandler struct {
	service service.UserAdminService
}

func NewAdminHandler(service service.UserAdminService) AdminHandler {
	return AdminHandler{service: service}
}

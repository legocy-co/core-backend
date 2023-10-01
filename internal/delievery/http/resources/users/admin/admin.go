package admin

import (
	"legocy-go/internal/domain/users/models/admin"
)

type AdminRegistrationResponse struct {
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Username string `json:"username"`
}

type AdminRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uar *AdminRegistrationRequest) ToAdmin() *admin.UserAdmin {
	return &admin.UserAdmin{
		Email:    uar.Email,
		Username: uar.Username,
		Role:     admin.ADMIN,
	}
}

func GetAdminResponse(ua *admin.UserAdmin) *AdminRegistrationResponse {
	return &AdminRegistrationResponse{
		Email:    ua.Email,
		Role:     ua.Role,
		Username: ua.Username,
	}
}

type UserAdminDetailResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

func GetUserAdminDetailResponse(ua *admin.UserAdmin) UserAdminDetailResponse {
	return UserAdminDetailResponse{
		ID:       ua.ID,
		Username: ua.Username,
		Email:    ua.Email,
		Role:     ua.Role,
	}
}

package admin

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type AdminRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (uar *AdminRegistrationRequest) ToAdmin() *models.UserAdminValueObject {
	return &models.UserAdminValueObject{
		Email:    uar.Email,
		Username: uar.Username,
		Role:     models.ADMIN,
	}
}

type AdminRegistrationResponse struct {
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Username string `json:"username"`
}

func GetAdminResponse(ua *models.UserAdminValueObject) *AdminRegistrationResponse {
	return &AdminRegistrationResponse{
		Email:    ua.Email,
		Role:     ua.Role,
		Username: ua.Username,
	}
}

type UserAdminUpdateRequest struct {
	models.UserAdminValueObject
}

func (r UserAdminUpdateRequest) ToUserAdminValueObject() (
	*models.UserAdminValueObject, error) {

	return &models.UserAdminValueObject{
		Username: r.Username,
		Email:    r.Email,
		Role:     r.Role,
	}, nil
}

type UserAdminDetailResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
}

func GetUserAdminDetailResponse(ua *models.UserAdmin) UserAdminDetailResponse {
	return UserAdminDetailResponse{
		ID:       ua.ID,
		Username: ua.Username,
		Email:    ua.Email,
		Role:     ua.Role,
	}
}

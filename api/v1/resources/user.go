package v1

import (
	models "legocy-go/pkg/auth/models"
)

type UserRegistrationResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ur *UserRegistrationRequest) ToUser() *models.User {
	return &models.User{
		Email:    ur.Email,
		Username: ur.Username,
		Role:     models.USER,
	}
}

func (ur *UserRegistrationRequest) ToAdmin() *models.User {
	return &models.User{
		Email:    ur.Email,
		Username: ur.Username,
		Role:     models.ADMIN,
	}
}

func GetUserResponse(u *models.User) *UserRegistrationResponse {
	return &UserRegistrationResponse{
		Email:    u.Email,
		Username: u.Username,
	}
}

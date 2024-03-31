package users

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FromDomain(user *models.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func FromDomainAdmin(user *models.UserAdmin) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

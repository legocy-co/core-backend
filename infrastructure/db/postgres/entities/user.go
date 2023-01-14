package postgres

import (
	models "legocy-go/pkg/auth/models"
)

type UserPostgres struct {
	Model
	Username string `gorm:"not null"`
	Email    string `gorm:"unique not null"`
	Password string
}

func FromUser(u *models.User, password string) *UserPostgres {
	return &UserPostgres{
		Username: u.Username,
		Email:    u.Email,
		Password: password,
	}
}

func (up *UserPostgres) ToUser() *models.User {
	return &models.User{
		Username: up.Username,
		Email:    up.Email,
	}
}

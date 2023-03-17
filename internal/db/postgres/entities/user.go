package postgres

import (
	models "legocy-go/internal/domain/auth/models"
)

type UserPostgres struct {
	Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Role     int
	Password string
}

func FromUser(u *models.User, password string) *UserPostgres {
	return &UserPostgres{
		Username: u.Username,
		Email:    u.Email,
		Password: password,
		Role:     u.Role,
	}
}

func (up *UserPostgres) ToUser() *models.User {
	return &models.User{
		ID:       int(up.ID),
		Username: up.Username,
		Email:    up.Email,
		Role:     up.Role,
	}
}

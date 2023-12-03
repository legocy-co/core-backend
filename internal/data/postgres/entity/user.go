package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserPostgres struct {
	Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Role     int
	Password string
}

func (up UserPostgres) TableName() string {
	return "users"
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

func (up *UserPostgres) GetUpdatedUserAdmin(
	vo models.UserAdminValueObject) *UserPostgres {
	up.Username = vo.Username
	up.Email = vo.Email
	up.Role = vo.Role
	return up
}

func FromUserAdminValueObject(vo models.UserAdminValueObject) *UserPostgres {
	return &UserPostgres{
		Username: vo.Username,
		Email:    vo.Email,
		Role:     vo.Role,
	}
}

func FromAdmin(u *models.UserAdmin, password string) *UserPostgres {
	return &UserPostgres{
		Username: u.Username,
		Email:    u.Email,
		Password: password,
		Role:     u.Role,
	}
}

func (up *UserPostgres) ToUserAdmin() *models.UserAdmin {
	return &models.UserAdmin{
		ID:       int(up.ID),
		Username: up.Username,
		Email:    up.Email,
		Role:     up.Role,
	}
}

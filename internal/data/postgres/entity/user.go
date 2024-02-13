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
	Images   []UserImagePostgres `gorm:"foreignKey:UserID"`
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

	if up.Images == nil {
		return models.NewUser(
			int(up.ID),
			up.Username,
			up.Email,
			up.Role,
			[]*models.UserImage{},
		)
	}

	images := make([]*models.UserImage, 0, len(up.Images))
	for _, img := range up.Images {
		images = append(images, img.ToUserImage())
	}

	return models.NewUser(
		int(up.ID),
		up.Username,
		up.Email,
		up.Role,
		images,
	)
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

func GetUpdatedUserEntity(vo models.UserValueObject, up *UserPostgres) *UserPostgres {
	up.Username = vo.Username
	up.Email = vo.Email
	return up
}

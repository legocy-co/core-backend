package postgres

import (
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

type UserPostgres struct {
	Model
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Role       int
	Password   string
	GoogleID   *string             `gorm:"unique;default:null"`
	FacebookID *string             `gorm:"unique;default:null"`
	Images     []UserImagePostgres `gorm:"foreignKey:UserID"`
}

func (up UserPostgres) TableName() string {
	return "users"
}

func FromUser(u *models.User, password string) *UserPostgres {
	return &UserPostgres{
		Username:   u.Username,
		Email:      u.Email,
		Password:   password,
		Role:       u.Role,
		GoogleID:   u.GoogleID,
		FacebookID: u.FacebookID,
	}
}

func FromVO(vo models.UserValueObject) *UserPostgres {
	return &UserPostgres{
		Username:   vo.Username,
		Email:      vo.Email,
		Role:       models.USER,
		GoogleID:   vo.GoogleID,
		FacebookID: vo.FacebookID,
	}
}

func (up *UserPostgres) ToUser() *models.User {

	if up.Images == nil {
		return models.NewUser(
			int(up.ID),
			up.Username,
			up.Email,
			up.Role,
			up.GoogleID,
			up.FacebookID,
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
		up.GoogleID,
		up.FacebookID,
		images,
	)
}

func (up *UserPostgres) GetUpdatedUserAdmin(
	vo models.UserAdminValueObject) *UserPostgres {
	up.Username = vo.Username
	up.Email = vo.Email
	up.Role = vo.Role
	up.GoogleID = vo.GoogleID
	up.FacebookID = vo.FacebookID
	return up
}

func FromAdminVO(u *models.UserAdminValueObject, password string) *UserPostgres {
	return &UserPostgres{
		Username:   u.Username,
		Email:      u.Email,
		Password:   password,
		Role:       u.Role,
		GoogleID:   u.GoogleID,
		FacebookID: u.FacebookID,
	}
}

func (up *UserPostgres) ToUserAdmin() *models.UserAdmin {
	return &models.UserAdmin{
		ID:         int(up.ID),
		Username:   up.Username,
		Email:      up.Email,
		Role:       up.Role,
		GoogleID:   up.GoogleID,
		FacebookID: up.FacebookID,
	}
}

func GetUpdatedUserEntity(vo models.UserValueObject, up *UserPostgres) *UserPostgres {
	up.Username = vo.Username
	up.Email = vo.Email
	up.GoogleID = vo.GoogleID
	up.FacebookID = vo.FacebookID
	return up
}

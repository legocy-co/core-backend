package auth

const (
	USER = iota
	ADMIN
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     int // Admin/User/etc.
	Images   []*UserImage
}

func NewUser(ID int, username string, email string, role int, images []*UserImage) *User {
	return &User{
		ID:       ID,
		Username: username,
		Email:    email,
		Role:     role,
		Images:   images,
	}
}

type UserValueObject struct {
	Username string
	Email    string
}

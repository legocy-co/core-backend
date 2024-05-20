package models

const (
	USER = iota
	ADMIN
)

type User struct {
	ID         int
	Username   string
	Email      string
	Role       int // Admin/User/etc.
	GoogleID   *string
	FacebookID *string
	Images     []*UserImage
}

func NewUser(ID int, username string, email string, role int, googleId, facebookId *string, images []*UserImage) *User {
	return &User{
		ID:         ID,
		Username:   username,
		Email:      email,
		Role:       role,
		Images:     images,
		GoogleID:   googleId,
		FacebookID: facebookId,
	}
}

type UserValueObject struct {
	Username   string
	Email      string
	GoogleID   *string
	FacebookID *string
}

// Only for user registration.
// TODO: delete me later: rewrite repository.Create to use ValueObject instead of full model
func FromVO(vo UserValueObject) *User {
	return &User{
		Username:   vo.Username,
		Email:      vo.Email,
		Role:       USER,
		GoogleID:   vo.GoogleID,
		FacebookID: vo.FacebookID,
	}
}

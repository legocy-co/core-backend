package auth

import (
	e "legocy-go/pkg/auth/errors"
)

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) Validate() error {
	if u.ID < 0 {
		return e.ErrUserAlreadyExists
	}

	return nil
}

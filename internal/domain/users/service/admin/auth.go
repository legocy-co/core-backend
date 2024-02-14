package admin

import (
	"context"
	"github.com/legocy-co/legocy/internal/app/errors"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
)

func (s UserAdminService) LoginAdmin(c context.Context, email string, password string) (*models.UserAdmin, *errors.AppError) {

	err := s.repo.ValidateUser(c, email, password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(c, email)

	return user, nil
}

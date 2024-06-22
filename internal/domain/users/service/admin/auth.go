package admin

import (
	"context"
	models "github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/internal/pkg/errors"
)

func (s UserAdminService) LoginAdmin(c context.Context, email string, password string) (*models.UserAdmin, *errors.AppError) {

	err := s.repo.ValidateUser(c, email, password)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(c, email)

	return user, nil
}

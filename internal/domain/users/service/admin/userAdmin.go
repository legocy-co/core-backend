package admin

import (
	"context"
	"legocy-go/internal/domain/users/models/admin"
	r "legocy-go/internal/domain/users/repository"
)

type UserAdminService struct {
	repo r.UserAdminRepository
}

func NewUserAdminService(repo r.UserAdminRepository) UserAdminService {
	return UserAdminService{repo: repo}
}

func (s *UserAdminService) CreateAdmin(
	c context.Context, u *admin.UserAdmin, password string) error {
	return s.repo.CreateAdmin(c, u, password)
}

func (s *UserAdminService) GetUserByEmail(
	c context.Context, email string) (*admin.UserAdmin, error) {
	return s.repo.GetUserByEmail(c, email)
}

func (s *UserAdminService) GetUserByID(c context.Context, id int) (*admin.UserAdmin, error) {
	return s.repo.GetUserByID(c, id)
}

func (s *UserAdminService) DeleteUser(c context.Context, id int) error {
	return s.repo.DeleteUser(c, id)
}

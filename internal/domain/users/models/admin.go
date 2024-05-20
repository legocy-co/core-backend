package models

type UserAdmin struct {
	ID         int
	Username   string
	Email      string
	Role       int // Admin/User/etc.
	GoogleID   *string
	FacebookID *string
}

type UserAdminValueObject struct {
	Username   string
	Email      string
	Role       int // Admin/User/etc.
	GoogleID   *string
	FacebookID *string
}

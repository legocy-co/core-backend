package auth

type UserAdmin struct {
	ID       int
	Username string
	Email    string
	Role     int // Admin/User/etc.
}

type UserAdminValueObject struct {
	Username string
	Email    string
	Role     int // Admin/User/etc.
}

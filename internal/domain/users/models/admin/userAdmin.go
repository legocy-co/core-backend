package admin

const (
	USER = iota
	ADMIN
)

type UserAdmin struct {
	ID       int
	Username string
	Email    string
	Role     int // Admin/User/etc.
}

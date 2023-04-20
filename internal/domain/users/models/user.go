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
}

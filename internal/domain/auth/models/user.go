package auth

const (
	USER = iota
	ADMIN
)

type User struct {
	ID       int
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     int    // Admin/User/etc.
}

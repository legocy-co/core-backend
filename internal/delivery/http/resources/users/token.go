package users

type JWTRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

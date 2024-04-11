package users

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type JWTResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

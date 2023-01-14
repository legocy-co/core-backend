package v1

type JWTRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTResponse struct {
	AccessToken string `json:"access"`
}

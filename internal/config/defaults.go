package config

var DefaultJWTConfig *JWTConfig = &JWTConfig{
	SecretKey:          "test12345",
	AccesTokenLifeTime: 3}

var DefaultMinioConfig *MinioConfig = &MinioConfig{
	Url:      "",
	User:     "",
	Password: "",
	Token:    "",
	Ssl:      false,
}

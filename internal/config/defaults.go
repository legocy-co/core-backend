package config

var DefaultJWTConfig *JWTConfig = &JWTConfig{
	SecretKey:          "test12345",
	AccesTokenLifeTime: 3,
}

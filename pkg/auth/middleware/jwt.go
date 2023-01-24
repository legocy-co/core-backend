package auth

import (
	"errors"
	c "legocy-go/config"
	auth "legocy-go/pkg/auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtConf *c.JWTConfig = c.GetJWTConfig()
var jwtKey string = JwtConf.SecretKey

type JWTClaim struct {
	Email string `json:"email"`
	Role  int    `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email string, role int) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Duration(JwtConf.AccesTokenLifeTime) * time.Hour)
	claims := &JWTClaim{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))
	return
}

func ParseTokenClaims(signedToken string) (*JWTClaim, bool) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, false
	}

	claims, ok := token.Claims.(*JWTClaim)
	return claims, ok
}

func ValidateToken(signedToken string) (err error) {
	claims, ok := ParseTokenClaims(signedToken)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}

func ValidateAdminToken(signedToken string) error {
	if err := ValidateToken(signedToken); err != nil {
		return err
	}

	claims, ok := ParseTokenClaims(signedToken)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.Role != auth.ADMIN {
		return errors.New("user is not admin")
	}

	return nil
}

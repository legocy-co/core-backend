package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

type JWTClaim struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  int    `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToken(
	email string,
	id, role int,
	secretKey string,
	lifetimeHours int,
) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Duration(lifetimeHours) * time.Hour)
	claims := &JWTClaim{
		ID:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secretKey))
	return
}

func GenerateRefreshToken(
	email string,
	id, role int,
	secretKey string,
	lifetimeHours int,
) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Duration(lifetimeHours) * time.Hour)
	claims := &JWTClaim{
		ID:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(secretKey))
	return
}

func ParseTokenClaims(signedToken string, secretKey string) (*JWTClaim, bool) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)

	if err != nil {
		log.Error(err.Error())
		return nil, false
	}

	claims, ok := token.Claims.(*JWTClaim)
	return claims, ok
}

func ValidateAccessToken(signedToken string, secretKey string) (err error) {
	claims, ok := ParseTokenClaims(signedToken, secretKey)
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

func ValidateAdminAccessToken(signedToken string, roleValue int, secretKey string) (err error) {

	claims, ok := ParseTokenClaims(signedToken, secretKey)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	if claims.Role != roleValue {
		err = errors.New("user is not admin")
		return
	}

	return
}

func ValidateRefreshToken(signedToken string, secretKey string) (err error) {
	claims, ok := ParseTokenClaims(signedToken, secretKey)
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

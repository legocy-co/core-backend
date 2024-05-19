package users

import (
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"google.golang.org/api/idtoken"
)

func FromGoogleToken(token *idtoken.Payload) models.UserValueObject {
	return models.UserValueObject{
		Email:    token.Claims["email"].(string),
		Username: token.Claims["name"].(string),
		GoogleID: &token.Subject,
	}
}

package users

import (
	"github.com/legocy-co/legocy/internal/domain/users/models"
	"github.com/legocy-co/legocy/pkg/facebook"
)

type FacebookSignInUpRequest struct {
	Token string `json:"token"`
}

func FromFacebookPayload(payload facebook.TokenPayload) models.UserValueObject {
	return models.UserValueObject{
		Email:      payload.Email,
		Username:   payload.Name,
		FacebookID: &payload.ID,
	}
}

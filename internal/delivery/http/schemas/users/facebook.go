package users

import "github.com/legocy-co/legocy/internal/domain/users/models"

type FacebookSignUpRequest struct {
	FacebookID string `json:"facebook_id" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Username   string `json:"username" binding:"required"`
}

func (r FacebookSignUpRequest) ToVO() models.UserValueObject {
	return models.UserValueObject{
		Username:   r.Username,
		Email:      r.Email,
		FacebookID: &r.FacebookID,
	}
}

type FacebookSignInRequest struct {
	FacebookID string `json:"facebook_id" binding:"required"`
}

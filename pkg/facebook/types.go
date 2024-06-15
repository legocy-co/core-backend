package facebook

import "github.com/legocy-co/legocy/internal/domain/users/models"

type TokenPayload struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (t *TokenPayload) ToUserVO() models.UserValueObject {
	return models.UserValueObject{
		Email:      t.Email,
		Username:   t.Name,
		FacebookID: &t.ID,
	}
}

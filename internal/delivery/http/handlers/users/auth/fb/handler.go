package fb

import (
	"github.com/legocy-co/legocy/internal/domain/users/repository"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

type Handler struct {
	r repository.UserExternalAuthRepository
}

func NewHandler(app *app.App) Handler {
	return Handler{
		r: app.GetFacebookAuthRepository(),
	}
}

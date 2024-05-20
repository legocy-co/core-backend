package google

import (
	"github.com/legocy-co/legocy/internal/domain/users/repository"
)

type Handler struct {
	r repository.UserExternalAuthRepository
}

func NewHandler(r repository.UserExternalAuthRepository) Handler {
	return Handler{r: r}
}

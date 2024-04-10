package like

import (
	marketplace "github.com/legocy-co/legocy/internal/domain/marketplace/repository"
)

type Handler struct {
	r marketplace.LikeRepository
}

func NewHandler(r marketplace.LikeRepository) *Handler {
	return &Handler{
		r: r,
	}
}

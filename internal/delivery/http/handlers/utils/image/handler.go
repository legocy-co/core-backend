package image

import "github.com/legocy-co/legocy/pkg/storage/client"

type Handler struct {
	storage client.ImageStorage
}

func NewHandler(storage client.ImageStorage) *Handler {
	return &Handler{storage: storage}
}

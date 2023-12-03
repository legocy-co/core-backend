package user_collection

import service "github.com/legocy-co/legocy/internal/domain/collections/service/collection"

type UserLegoCollectionHandler struct {
	s service.UserCollectionService
}

func NewUserLegoCollectionHandler(s service.UserCollectionService) UserLegoCollectionHandler {
	return UserLegoCollectionHandler{s: s}
}

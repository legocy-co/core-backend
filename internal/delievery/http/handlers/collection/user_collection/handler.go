package user_collection

import service "legocy-go/internal/domain/collections/service/collection"

type UserLegoCollectionHandler struct {
	s service.UserCollectionService
}

func NewUserLegoCollectionHandler(s service.UserCollectionService) UserLegoCollectionHandler {
	return UserLegoCollectionHandler{s: s}
}

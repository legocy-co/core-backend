package admin

import (
	s "github.com/legocy-co/legocy/internal/domain/calculator/service/admin"
)

type Handler struct {
	service s.LegoSetValuationAdminService
}

func NewHandler(service s.LegoSetValuationAdminService) Handler {
	return Handler{service: service}
}

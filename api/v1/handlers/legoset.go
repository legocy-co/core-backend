package v1

import (
	"legocy-go/pkg/lego/usecase"
	"net/http"
)

func ListSet(service usecase.LegoSetUseCase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// todo:
	})
}

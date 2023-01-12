package v1

import (
	"legocy-go/pkg/lego/usecase"
	"net/http"
)

func ListSeries(service usecase.LegoSeriesUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("TODO"))
	})
}

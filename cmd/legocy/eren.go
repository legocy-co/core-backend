package main

import (
	router "github.com/legocy-co/legocy/internal/delivery/http/router/v1"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func main() {
	_app := app.New()

	srv := router.GetV1Router(_app)
	if err := srv.Run("8080"); err != nil {
		panic(err)
	}
}

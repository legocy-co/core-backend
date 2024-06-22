package main

import (
	"github.com/legocy-co/legocy/internal/delivery/http/server"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func main() {
	_app, err := app.New()
	if err != nil {
		panic(err)
	}

	srv := server.New(_app)
	if err := srv.Run("8080"); err != nil {
		panic(err)
	}
}

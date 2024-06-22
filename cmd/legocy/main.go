package main

import (
	"github.com/legocy-co/legocy/internal/delivery/http/server"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func main() {
	a := app.New()

	srv := server.New(a)
	if err := srv.Run("8080"); err != nil {
		panic(err)
	}
}

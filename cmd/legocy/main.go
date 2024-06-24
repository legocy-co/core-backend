package main

import (
	"github.com/legocy-co/legocy/internal/delivery/http/server"
	"github.com/legocy-co/legocy/internal/pkg/app"
	"log"
)

func main() {
	_app, err := app.New()
	if err != nil {
		log.Fatalf("Error while initializing application: %v", err)
	}

	srv := server.New(_app)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}

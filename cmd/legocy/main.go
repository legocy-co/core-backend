package main

import (
	"github.com/joho/godotenv"
	"legocy-go/internal/app"
	r "legocy-go/internal/delivery/http/router/v1"
)

func main() {
	godotenv.Load()

	_app := app.New()

	v1 := r.GetV1Router(_app)
	v1.Run("8080")
}

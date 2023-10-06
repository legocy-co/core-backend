package main

import (
	"github.com/joho/godotenv"
	"legocy-go/internal/app"
	r "legocy-go/internal/delievery/http/router"
)

func main() {
	godotenv.Load()

	_app := app.New()

	v1 := r.InitRouter(_app)
	v1.Run("8080")
}

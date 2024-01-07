package main

import (
	"github.com/legocy-co/legocy/internal/app"
	r "github.com/legocy-co/legocy/internal/delivery/http/router/v1"
)

func main() {
	_app := app.New()

	v1 := r.GetV1Router(_app)
	v1.Run("8080")
}

package main

import (
	r "github.com/legocy-co/legocy/internal/delivery/http/router/v1"
	"github.com/legocy-co/legocy/internal/pkg/app"
)

func main() {
	_app := app.New()

	v1 := r.GetV1Router(_app)
	v1.Run("8080")
}

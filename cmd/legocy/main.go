package main

import (
	"fmt"
	"legocy-go/internal/app"
	r "legocy-go/internal/delievery/http/router"
	"legocy-go/pkg/helpers"
)

var configFilepath string = helpers.GetConfigFilepath("/internal/config/json/config.json")

func main() {
	fmt.Printf("config fp: %v", configFilepath)
	_app := app.New(configFilepath)

	v1 := r.InitRouter(_app)
	v1.Run("8080")
}

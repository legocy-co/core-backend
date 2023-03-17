package main

import (
	"legocy-go/internal/app"
	r "legocy-go/internal/delievery/http/router"
	"os"
)

const configFilepath = "/internal/config/json/config.json"

func main() {
	cwd, _ := os.Getwd()
	_app := app.New(cwd + configFilepath)

	v1 := r.InitRouter(_app)
	v1.Run("8080")
}

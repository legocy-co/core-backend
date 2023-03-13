package main

import (
	r "legocy-go/delievery/http/router"
	"legocy-go/internal/app"
	"os"
)

const configFilepath = "/internal/config/json/config.json"

func main() {
	cwd, _ := os.Getwd()
	_app := app.New(cwd + configFilepath)

	v1 := r.InitRouter(_app)
	v1.Run("8080")
}

package main

import (
	r "legocy-go/delievery/http/router"
	"legocy-go/internal/app"
)

const configFilepath = "/Users/wjojf/GolandProjects/legocy-go-clean/internal/config/json/config.json"

func main() {
	_app := app.New(configFilepath)

	v1 := r.InitRouter(_app)
	v1.Run("8080")
}

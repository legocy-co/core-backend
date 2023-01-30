package main

import (
	r "legocy-go/api/v1/router"
	"legocy-go/internal/app"
)

const configFilepath = "/Users/wjojf/GolandProjects/legocy-go-clean/internal/config/json/config.json"

func main() {
	app := app.New(configFilepath)
	router := r.InitRouter(app)
	router.Run("8080")
}

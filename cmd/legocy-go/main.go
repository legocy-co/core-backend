package main

import "legocy-go/internal/pkg/app"

const configFilepath = "/Users/wjojf/GolandProjects/legocy-go-clean/internal/config/json/config.json"

func main() {
	app := app.New(configFilepath)
	app.Run("8080")
}

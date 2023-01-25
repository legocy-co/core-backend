package main

import "legocy-go/internal/pkg/app"

func main() {
	app := app.New("/Users/wjojf/GolandProjects/legocy-go-clean/config/json/config.json")
	app.Run("8080")
}

package main

import (
	"legocy-go/config"
	"legocy-go/server"
)

func main() {

	// TODO: Write real logic
	db := config.POSTGRES_DB

	if db == "" {
		return
	}

	app := server.NewApp()
	app.Run()
}

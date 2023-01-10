package main

import (
	"legocy-go/config"
	"legocy-go/server"
	"log"
)

func main() {

	if err := config.Init(); err != nil {
		log.Fatalln("Could not Read config")
		return
	}

	app := server.NewApp()
	app.Run()
}

package main

import (
	rv1 "legocy-go/api/v1/router"
)

func main() {
	router := rv1.InitRouter()
	router.Run(":" + "8080")
}

package main

import (
	v1_routers "legocy-go/routers/api/v1"
	"strconv"
)

func main() {
	port := 8080
	r := v1_routers.SetupRouter()
	r.Run(":" + strconv.Itoa(port))
}

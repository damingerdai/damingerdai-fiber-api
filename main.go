package main

import (
	"github.com/dammingerdai/damingerdai-fiber-api/internal/routers"
)

func main() {
	app := routers.NewRouter()

	app.Listen(":3000")

}

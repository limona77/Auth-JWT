package main

import (
	"auth/internal/app"
)

const configPath = "config/config.yaml"

func main() {

	app.Run(configPath)

}

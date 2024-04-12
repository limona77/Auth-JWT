package main

import (
	"auth/internal/app"
)

const configPath = "config/config.yaml"

// @title Auth-JWT
// @version 1.0
// @description Auth-JWT
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description					JWT token
func main() {
	app.Run(configPath)
}

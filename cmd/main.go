package main

import (
	"auth/internal/app"
)

// @title Auth-JWT
// @version 1.0
// @description Auth-JWT
// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description					JWT token

// @securityDefinitions.apikey  Cookie
// @in                          header
// @name                        refreshToken
// @description					Refresh token

func main() {
	app.Run()
}

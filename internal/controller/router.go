package controller

import (
	"auth/internal/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, services *service.Services) {
	auth := app.Group("/auth")
	newAuthRoutes(auth, services.Auth)
	uR := &clientRoutes{clientService: services.Client}
	app.Get("/me", uR.getUser)
}

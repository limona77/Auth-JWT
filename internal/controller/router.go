package controller

import (
	"auth/internal/service"

	"github.com/gofiber/swagger"

	_ "auth/docs"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, services *service.Services) {
	auth := app.Group("/auth")
	newAuthRoutes(auth, services.Auth)
	uR := &clientRoutes{clientService: services.Client}
	app.Get("/me", uR.getUser)
	app.Get("/swagger/*", swagger.HandlerDefault)
}

package controller

import (
	"auth/internal/service"
	"github.com/gofiber/fiber/v2"
)

type authRoutes struct {
	authService service.Auth
}

func newAuthRoutes(g fiber.Router, authService service.Auth) {
	aR := &authRoutes{authService: authService}
	g.Post("/register", aR.register)

}

func (aR *authRoutes) register(c *fiber.Ctx) error {
	return nil
}

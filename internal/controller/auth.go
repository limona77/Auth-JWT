package controller

import (
	custom_validator "auth/internal/custom-validator"
	"auth/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

type authRoutes struct {
	authService service.Auth
}

func newAuthRoutes(g fiber.Router, authService service.Auth) {
	aR := &authRoutes{authService: authService}
	g.Post("/register", aR.register)

}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (aR *authRoutes) register(c *fiber.Ctx) error {
	path := "internal.controller.auth.register"

	var uC UserCredentials

	err := c.BodyParser(&uC)
	if err != nil {
		return err
	}

	v := &custom_validator.XValidator{}

	err = v.Validate(uC)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".Validate, error: {%w}", err).Error())
		wrapHttpError(c, fiber.StatusBadRequest, err.Error())
	}

	return nil
}

package controller

import (
	custom_errros "auth/internal/custom-errros"
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
	Email    string `json:"email" validate:"required,email,min=8"`
	Password string `json:"password" validate:"required,min=5"`
}

func (aR *authRoutes) register(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.register"

	var uC UserCredentials

	err := ctx.BodyParser(&uC)
	if err != nil {
		return err
	}

	v := &custom_validator.XValidator{}

	err = v.Validate(uC)
	if err != nil {

		slog.Errorf(fmt.Errorf(path+".Validate, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusBadRequest, err.Error())
	}
	err = aR.authService.CreateUser(ctx.Context(), service.AuthParams{uC.Email, uC.Password})
	if err != nil {
		if err == custom_errros.ErrAlreadyExists {
			slog.Errorf(fmt.Errorf(path+".CreateUser, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, err.Error())
		}
		slog.Errorf(fmt.Errorf(path+".CreateUser, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

package controller

import (
	"auth/internal/custom-errors"
	custom_validator "auth/internal/custom-validatot"
	"auth/internal/service"
	"errors"
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

	v := &custom_validator.XValidator{Validator: custom_validator.Validate}

	err = v.Validate(uC)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".Validate, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusBadRequest, err.Error())
	}
	authParams := service.AuthParams{Email: uC.Email, Password: uC.Password}

	user, err := aR.authService.CreateUser(ctx.Context(), authParams)
	if err != nil {
		if errors.Is(err, custom_errors.ErrAlreadyExists) {
			slog.Errorf(fmt.Errorf(path+".CreateUser, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, err.Error())
		}
		slog.Errorf(fmt.Errorf(path+".CreateUser, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	fmt.Println("jopa", user)
	tokens, err := aR.authService.GenerateTokens(ctx.Context(), authParams)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})

	resp := map[string]interface{}{"user": user, "refreshToken": tokens.RefreshToken, "accessToken": tokens.AccessToken}
	err = httpResponse(ctx, 200, resp)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".JSON, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

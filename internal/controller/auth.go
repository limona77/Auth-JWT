package controller

import (
	"auth/internal/custom-errors"
	custom_validator "auth/internal/custom-validator"
	"auth/internal/model"
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
	g.Post("/login", aR.login)
	g.Get("/refresh", aR.refresh)

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
		slog.Errorf(fmt.Errorf(path+".BodyParser, error: {%w}", err).Error())
		return wrapHttpError(ctx, 500, "internal error")
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
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrAlreadyExists.Error())
		}
		slog.Errorf(fmt.Errorf(path+".CreateUser, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	tokens, _, err := aR.authService.GenerateTokens(ctx.Context(), authParams)
	if err != nil {
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}

	tokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	token, err := aR.authService.SaveToken(ctx.Context(), tokenModel)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
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

func (aR *authRoutes) login(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.login"

	var uC UserCredentials
	err := ctx.BodyParser(&uC)
	if err != nil {
		return wrapHttpError(ctx, 500, "internal error")
	}
	v := &custom_validator.XValidator{Validator: custom_validator.Validate}

	err = v.Validate(uC)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".Validate, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusBadRequest, err.Error())
	}

	authParams := service.AuthParams{Email: uC.Email, Password: uC.Password}

	tokens, user, err := aR.authService.GenerateTokens(ctx.Context(), authParams)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrUserNotFound.Error())
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrWrongCredetianls.Error())
		}
		slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	tokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	token, err := aR.authService.SaveToken(ctx.Context(), tokenModel)
	if err != nil {
		return wrapHttpError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
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
func (aR *authRoutes) refresh(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.refresh"

	refreshToken := ctx.Cookies("refreshToken")

	if refreshToken == "" {
		return wrapHttpError(ctx, fiber.StatusBadRequest, "refresh token is required")
	}
	c := &service.ClientService{}
	tokenClaims, err := c.VerifyToken(refreshToken)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".VerifyToken, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusUnauthorized, custom_errors.ErrUserUnauthorized.Error())
	}

	tokenModel := model.Token{UserID: tokenClaims.UserID, RefreshToken: refreshToken}
	_, err = aR.authService.GetToken(ctx.Context(), tokenModel)
	if err != nil {
		return err
	}

	authParams := service.AuthParams{Email: tokenClaims.Email}
	tokens, user, err := aR.authService.GenerateTokens(ctx.Context(), authParams)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrUserNotFound.Error())
		}
		if errors.Is(err, custom_errors.ErrWrongCredetianls) {
			slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrWrongCredetianls.Error())
		}
		slog.Errorf(fmt.Errorf(path+".GetUserByEmail, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	newRefreshTokenModel := model.Token{
		RefreshToken: tokens.RefreshToken,
		UserID:       user.ID,
	}
	token, err := aR.authService.SaveToken(ctx.Context(), newRefreshTokenModel)
	if err != nil {
		return wrapHttpError(ctx, fiber.StatusInternalServerError, err.Error())
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    token.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})
	return nil
}

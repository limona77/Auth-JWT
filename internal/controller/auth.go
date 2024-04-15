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
	g.Get("/logout", aR.logout)
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email,min=8"`
	Password string `json:"password" validate:"required,min=5"`
}

type authResponse struct {
	User         model.User `json:"user"`
	RefreshToken string     `json:"refreshToken"`
	AccessToken  string     `json:"accessToken"`
}

// @Summary Register
// @Tags auth
// @Description user registration
// @Accept json
// @Produce json
// @ID register
// @Param input body UserCredentials true "credentials"
// @Success 200 {object} authResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
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

	tokens, user, err := aR.authService.Register(ctx.Context(), authParams)
	if err != nil {
		if errors.Is(err, custom_errors.ErrAlreadyExists) {
			slog.Errorf(fmt.Errorf(path+".Register, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrAlreadyExists.Error())
		}
		slog.Errorf(fmt.Errorf(path+".Register, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})
	user.Password = ""
	resp := authResponse{User: user, RefreshToken: tokens.RefreshToken, AccessToken: tokens.AccessToken}
	err = httpResponse(ctx, 200, resp)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".httpResponse, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

// @Summary Login
// @Tags auth
// @Description user login
// @Accept json
// @Produce json
// @ID login
// @Param input body UserCredentials true "credentials"
// @Success 200 {object} authResponse "ok"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
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

	tokens, user, err := aR.authService.Login(ctx.Context(), authParams)
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

	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})

	resp := authResponse{User: user, RefreshToken: tokens.RefreshToken, AccessToken: tokens.AccessToken}
	err = httpResponse(ctx, 200, resp)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".httpResponse, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

// @Summary Refresh
// @Security Cookie
// @Tags auth
// @Description user refresh
// @Accept json
// @Produce json
// @ID refresh
// @Success 200 {object} authResponse "ok"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/refresh [get]
func (aR *authRoutes) refresh(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.refresh"

	refreshToken := ctx.Cookies("refreshToken")
	if refreshToken == "" {
		slog.Errorf(fmt.Errorf(path+".Cookies, error: {%w}", errors.New("token is required")).Error())
		return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrUserUnauthorized.Error())
	}

	tokens, user, err := aR.authService.Refresh(ctx.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserUnauthorized) {
			slog.Errorf(fmt.Errorf(path+".Refresh, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusUnauthorized, custom_errors.ErrUserUnauthorized.Error())
		}
		slog.Errorf(fmt.Errorf(path+".Refresh, error: {%w}", err).Error())
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		MaxAge:   30 * 24 * 60 * 60 * 1000,
		HTTPOnly: true,
	})
	resp := authResponse{User: user, RefreshToken: tokens.RefreshToken, AccessToken: tokens.AccessToken}
	err = httpResponse(ctx, 200, resp)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".httpResponse, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

// @Summary Logout
// @Security Cookie
// @Tags auth
// @Description user logout
// @Accept json
// @Produce json
// @ID logout
// @Success 200 {object} int "ok"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/logout [get]
func (aR *authRoutes) logout(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refreshToken")
	path := "internal.controller.auth.logout"
	if refreshToken == "" {
		slog.Errorf(fmt.Errorf(path+".Cookies, error: {%w}", errors.New("token is required")).Error())
		return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrUserUnauthorized.Error())
	}
	userID, err := aR.authService.Logout(ctx.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserUnauthorized) {
			slog.Errorf(fmt.Errorf(path+".Logout, error: {%w}", err).Error())
			return wrapHttpError(ctx, fiber.StatusBadRequest, custom_errors.ErrUserUnauthorized.Error())
		}
		slog.Errorf(fmt.Errorf(path+".Logout, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
	})
	resp := map[string]interface{}{"userID": userID}
	err = httpResponse(ctx, 200, resp)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".httpResponse, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

package controller

import (
	custom_errors "auth/internal/custom-errors"
	"auth/internal/service"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
	"strings"
)

type clientRoutes struct {
	clientService service.Client
}

func (aR *clientRoutes) getUser(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.getUser"
	accessToken := ctx.Get("Authorization")
	if len([]rune(accessToken)) == 0 {
		return fmt.Errorf(path, "no token provided")
	}
	t := strings.Split(accessToken, " ")[1]
	email, err := aR.clientService.VerifyToken(t)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".VerifyToken, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusUnauthorized, custom_errors.ErrUserUnauthorized.Error())
	}
	params := service.AuthParams{Email: email}
	user, err := aR.clientService.GetUserByEmail(ctx.Context(), params)
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
	resp := map[string]interface{}{"user": user}
	err = httpResponse(ctx, fiber.StatusOK, resp)
	if err != nil {
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

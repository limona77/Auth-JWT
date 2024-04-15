package controller

import (
	custom_errors "auth/internal/custom-errors"
	"auth/internal/model"
	"auth/internal/service"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/slog"
)

type clientRoutes struct {
	clientService service.Client
}
type clientResponse struct {
	User model.User `json:"user"`
}

// @Summary AuthMe
// @Security JWT
// @Tags client
// @Description check auth
// @Accept json
// @Produce json
// @ID get-user
// @Success 200 {object} clientResponse "ok"
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /me [get]
func (aR *clientRoutes) authMe(ctx *fiber.Ctx) error {
	path := "internal.controller.auth.getUser"
	accessToken := ctx.Get("Authorization")
	fmt.Println("accessToken", accessToken)
	if len([]rune(accessToken)) == 0 {
		return fmt.Errorf(path, "no token provided")
	}
	t := strings.Split(accessToken, " ")[1]
	tokenClaims, err := aR.clientService.VerifyToken(t)
	if err != nil {
		slog.Errorf(fmt.Errorf(path+".VerifyToken, error: {%w}", err).Error())
		return wrapHttpError(ctx, fiber.StatusUnauthorized, custom_errors.ErrUserUnauthorized.Error())
	}
	params := service.AuthParams{Email: tokenClaims.Email}
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
	resp := clientResponse{User: user}
	err = httpResponse(ctx, fiber.StatusOK, resp)
	if err != nil {
		return wrapHttpError(ctx, fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}

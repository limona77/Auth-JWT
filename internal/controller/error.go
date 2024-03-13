package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func wrapHttpError(ctx *fiber.Ctx, code int, message string) error {
	err := errors.New(message)

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		return ctx.Status(code).SendString(message)
	}

	return httpResponse(ctx, code, fiber.Map{"message": message})
}

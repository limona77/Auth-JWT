package controller

import (
	"github.com/gofiber/fiber/v2"
)

func httpResponse(ctx *fiber.Ctx, code int, params interface{}) error {
	err := ctx.Status(code).JSON(params)
	if err != nil {
		return err
	}
	return nil
}

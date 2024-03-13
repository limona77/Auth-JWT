package controller

import (
	"github.com/gofiber/fiber/v2"
)

func httpResponse(ctx *fiber.Ctx, code int, params map[string]interface{}) error {
	response := fiber.Map{}
	for k, v := range params {
		response[k] = v
	}
	err := ctx.Status(code).JSON(response)
	if err != nil {
		return err
	}
	return nil
}

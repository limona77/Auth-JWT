package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func wrapHttpError(c *fiber.Ctx, errStatus int, message string) {
	err := errors.New(message)
	_, ok := err.(*fiber.Error)
	if !ok {
		report := fiber.NewError(errStatus, err.Error())
		_ = c.JSON(errStatus, report.Error())
	}
	_ = c.Status(errStatus).SendString(errors.New(message).Error())

}

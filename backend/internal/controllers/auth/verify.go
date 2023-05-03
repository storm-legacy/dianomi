package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

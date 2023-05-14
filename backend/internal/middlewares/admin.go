package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role != "administrator" {
		return c.SendStatus(fiber.StatusForbidden)
	}
	return c.Next()
}

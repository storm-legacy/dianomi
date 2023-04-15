package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/config"
)

func PublicKey(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"publicKey": config.GetString("APP_JWT_EDDSA_PUBLIC_KEY"),
	})
}

func Register(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

func Update(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

func Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

func Verify(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"status": "not implemented",
	})
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/config"
)

func PublicKey(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"publickey": config.GetString("APP_JWT_EDDSA_PUBLIC_KEY")})
}

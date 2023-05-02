package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/jwt"
)

func Logout(c *fiber.Ctx) error {
	jti := c.Locals("jti").(string)
	exp := c.Locals("exp").(time.Time)
	jwt.RevokeToken(jti, exp)

	log.WithField("jti", jti).Debug("Token was revoked")
	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data:   "User was successfuly logged out!",
	})
}

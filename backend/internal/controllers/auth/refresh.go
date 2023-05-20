package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/jwt"
)

func Refresh(c *fiber.Ctx) error {
	jti := c.Locals("jti").(string)
	exp := c.Locals("exp").(time.Time)
	sub := c.Locals("sub").(uint64)
	role := c.Locals("role").(string)
	verified := c.Locals("verified").(bool)

	jwt.RevokeToken(jti, exp)
	claims := make(map[string]interface{})
	claims["role"] = role
	claims["verfied"] = verified

	token, err := jwt.GenerateToken(sub, claims)
	if err != nil {
		log.WithField("err", err).Error("Problem while refreshing JWT token")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.WithField("sub", sub).Debug("Account refreshed it's token")
	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data: fiber.Map{
			"token": token,
		},
	})
}

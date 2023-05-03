package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/jwt"
)

func Minio(c *fiber.Ctx) error {
	token := c.FormValue("token")
	claims, err := jwt.ExtractClaims(token)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"reason": "Missing or broken token"})
	}

	// Check if user is admin
	role := (*claims)["role"].(string)
	if role != "administrator" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"reason": "User is not an administrator"})
	}

	// Check if token is valid
	now := time.Now().Unix()
	tokenSub := uint64((*claims)["sub"].(float64))
	tokenExp := time.Unix(int64((*claims)["exp"].(float64)), 0)
	tokenNbf := time.Unix(int64((*claims)["nbf"].(float64)), 0)
	tokenJti := (*claims)["jti"].(string)

	// Is token valid
	if now > tokenExp.Unix() ||
		now < tokenNbf.Unix() ||
		jwt.IsTokenRevoked(tokenJti) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"reason": "Token is invalid"})
	}

	minioClaims := make(map[string]string)
	minioClaims["bucket"] = "dianomi-videos"
	minioClaims["permissions"] = "writeonly"

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":               fmt.Sprintf("user-%d", tokenSub),
		"maxValiditySeconds": 3600,
		"claims":             minioClaims,
	})
}

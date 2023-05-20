package middlewares

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Extract header with token
	splitToken := strings.Split(c.Get("Authorization"), "Bearer ")
	// Check if token was provided
	if len(splitToken) < 2 {
		log.WithField("token", splitToken).Debug("User tried connection without token")
		return c.Status(fiber.StatusUnauthorized).JSON(mod.Response{
			Status: "error",
			Data:   "Missing authorization token",
		})
	}

	// Extract claims
	token := splitToken[1]
	claims, err := jwt.ExtractClaims(token)
	if err != nil {
		log.WithFields(log.Fields{
			"token": splitToken,
			"err":   err.Error(),
		}).Debug("Invalid authorization token")

		return c.Status(fiber.StatusUnauthorized).JSON(mod.Response{
			Status: "error",
			Data:   "Invalid authorization token",
		})
	}

	// Check if token is valid
	now := time.Now().Unix()
	tokenSub := uint64((*claims)["sub"].(float64))
	tokenExp := time.Unix(int64((*claims)["exp"].(float64)), 0)
	tokenNbf := time.Unix(int64((*claims)["nbf"].(float64)), 0)
	tokenJti := (*claims)["jti"].(string)
	verified := (*claims)["verified"].(bool)

	// Is user verified
	if !verified {
		log.WithFields(log.Fields{
			"verified": verified,
			"sub":      tokenSub,
		}).Debug("User is not verified")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Is token valid
	if now > tokenExp.Unix() ||
		now < tokenNbf.Unix() ||
		jwt.IsTokenRevoked(tokenJti) {
		log.WithField("jti", tokenJti).Debug("Usage of unauthorized/revoked token")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Extract additional data
	role := (*claims)["role"].(string)

	c.Locals("sub", tokenSub)
	c.Locals("role", role)
	c.Locals("jti", tokenJti)
	c.Locals("exp", tokenExp)
	c.Locals("verified", verified)

	return c.Next()
}

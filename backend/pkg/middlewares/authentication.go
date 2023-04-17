package middlewares

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/jwt"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func ValidateToken(c *fiber.Ctx) error {
	// Extract header with token
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "missing authorization token"})
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "missing authorization token"})
	}
	token := splitToken[1]

	// Parse token (will validate it's origin)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		log.WithField("err", err).Error("problem while parsing jwt token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	tokenExp := time.Unix(int64(claims["exp"].(float64)), 0)

	// Check if token expired
	if time.Now().Unix() > tokenExp.Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "token expired"})
	}

	// Check if token wasn't revoked
	// DON'T DO IT THIS WAY
	// Rather use something like Redis or another fast access cache
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("problem connection to database while checking token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	ctx := context.Background()
	queries := sqlc.New(db)
	defer db.Close()
	// Database block (end)

	// Check if token was revoked
	_, err = queries.CheckToken(ctx, token)
	if err != sql.ErrNoRows {
		if err != nil {
			log.WithField("err", err).Error("Error while checking revoked tokens")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "token was revoked"})
	}

	role := claims["role"].(string)
	c.Locals("role", role)
	c.Locals("token", token)
	return c.Next()
}

// Yeah, hard coding permissions always ends well
func PremiumMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")
	fmt.Print(role)
	if role == "premium" || role == "administrator" {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "not enough permissions"})
}

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")
	if role == "administrator" {
		return c.Next()
	}
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "not enough permissions"})
}

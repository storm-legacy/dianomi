package main

import (
	_ "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/storm-legacy/dianomi/cmd/dianomi-server/internal/routes"
	mw "github.com/storm-legacy/dianomi/pkg/middlewares"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Health
	app.Get("/api/v1/healthcheck", routes.Healthcheck)
	app.Get("/api/v1/_healthcheck", mw.WhitelistMiddleware, routes.InternalHealthcheck)

	// Auth
	// TODO refreshToken, check if token is present
	app.Post("/api/v1/auth/login", routes.Login)
	app.Post("/api/v1/auth/register", routes.Register)
	app.Post("/api/v1/auth/logout", mw.ValidateToken, routes.Logout)
	// app.Post("/api/v1/auth/account", mw.ValidateToken, routes.Update)
	// app.Get("/api/v1/auth/verify", mw.ValidateToken, routes.Verify)
	app.Get("/api/v1/auth/publickey", routes.PublicKey)

	// Testing routes
	app.Get("/api/v1/routeNormal", mw.ValidateToken, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Free user route"})
	})

	app.Get("/api/v1/routePremium", mw.ValidateToken, mw.PremiumMiddleware, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Premium user route"})
	})

	app.Get("/api/v1/routeAdmin", mw.ValidateToken, mw.AdminMiddleware, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Administrator route"})
	})

	log.Fatal(app.Listen(":3000"))
}

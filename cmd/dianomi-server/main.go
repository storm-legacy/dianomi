package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/storm-legacy/dianomi/cmd/dianomi-server/internal/routes"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(), func(c *fiber.Ctx) error { return c.Next() })

	// Auth
	// TODO new field in verification table with type of code
	// TODO add revocation table for JWT tokens
	app.Get("/api/v1/auth/publickey", routes.PublicKey)
	app.Post("/api/v1/auth/register", routes.Register)
	app.Get("/api/v1/auth/verify", routes.Verify)
	app.Post("/api/v1/auth/delete", routes.Delete)
	app.Post("/api/v1/auth/update", routes.Update)
	app.Post("/api/v1/auth/logout", routes.Logout)

	log.Fatal(app.Listen(":3000"))
}

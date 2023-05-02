package main

import (
	_ "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	ctrl "github.com/storm-legacy/dianomi/internal/controllers"
	authCtrl "github.com/storm-legacy/dianomi/internal/controllers/auth"
	mid "github.com/storm-legacy/dianomi/internal/middlewares"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api/v1")

	// * Healthcheck
	api.Get("/healthcheck", ctrl.Healthcheck)
	api.Get("/_healthcheck", ctrl.InternalHealthcheck)

	// * Authentication group
	auth := api.Group("auth")
	auth.Post("/publickey", authCtrl.PublicKey)
	auth.Post("/verify", ctrl.NotImplemented)
	auth.Post("/login", authCtrl.Login)
	auth.Post("/register", authCtrl.Register)
	auth.Post("/refresh", mid.AuthMiddleware, authCtrl.Refresh)
	auth.Post("/logout", mid.AuthMiddleware, authCtrl.Logout)

	// * User group
	user := api.Group("user", mid.AuthMiddleware)
	user.Get("/account", ctrl.NotImplemented)
	user.Get("/list", ctrl.NotImplemented)

	// * Administration
	admin := api.Group("admin", mid.AuthMiddleware, mid.AdminMiddleware)
	admin.Post("/upload", ctrl.NotImplemented)

	// TODO TO BE REMOVED
	api.Get("/routeNormal", mid.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Free user route"})
	})

	api.Get("/routePremium", mid.AuthMiddleware, func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if !(role == "administrator" || role == "premium") {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Premium user route"})
	})

	api.Get("/routeAdmin", mid.AuthMiddleware, func(c *fiber.Ctx) error {
		role := c.Locals("role").(string)
		if role != "administrator" {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Administrator route"})
	})

	log.Fatal(app.Listen(":3000"))
}

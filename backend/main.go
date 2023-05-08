package main

import (
	_ "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	ctrl "github.com/storm-legacy/dianomi/internal/controllers"
	adminVideoCtrl "github.com/storm-legacy/dianomi/internal/controllers/admin/video"
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
	// auth.Get("/publickey", authCtrl.PublicKey)
	auth.Post("/verify", mid.AuthMiddleware, authCtrl.Verify)
	auth.Post("/login", authCtrl.Login)
	auth.Post("/register", authCtrl.Register)
	auth.Post("/refresh", mid.AuthMiddleware, authCtrl.Refresh)
	auth.Post("/logout", mid.AuthMiddleware, authCtrl.Logout)

	// * MINIO S3
	auth.Post("/minio", authCtrl.Minio)
	auth.Head("/minio", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	// * User group
	user := api.Group("user", mid.AuthMiddleware)
	user.Get("/account", ctrl.NotImplemented)
	user.Get("/list", ctrl.NotImplemented)

	// * Administration
	admin := api.Group("admin", mid.AuthMiddleware, mid.AdminMiddleware)
	admin.Post("/upload", ctrl.NotImplemented)

	// * Videos (Admin)
	adminVideo := admin.Group("video")
	adminVideo.Post("/", adminVideoCtrl.PostVideo)

	log.Fatal(app.Listen(":3000"))
}

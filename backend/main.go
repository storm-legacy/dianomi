package main

import (
	_ "time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	ctrl "github.com/storm-legacy/dianomi/internal/controllers"
	authCtrl "github.com/storm-legacy/dianomi/internal/controllers/auth"
	develCtrl "github.com/storm-legacy/dianomi/internal/controllers/devel"
	profileCtrl "github.com/storm-legacy/dianomi/internal/controllers/profile"
	testsCtrl "github.com/storm-legacy/dianomi/internal/controllers/test"
	usersCtrl "github.com/storm-legacy/dianomi/internal/controllers/users"
	packagesCtrl "github.com/storm-legacy/dianomi/internal/controllers/users/packages"
	videoCtrl "github.com/storm-legacy/dianomi/internal/controllers/video"
	videoCategoryCtrl "github.com/storm-legacy/dianomi/internal/controllers/video/category"
	mid "github.com/storm-legacy/dianomi/internal/middlewares"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api/v1")
	api.Get("/", mid.AuthMiddleware, func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })

	// https://localhost/test
	api.Post("/test", testsCtrl.Test)

	// * DEVELOPMENT ENDPOINTS
	dev := api.Group("dev")
	dev.Post("/setpassword", develCtrl.SetPassword)
	dev.Post("/setpackage", develCtrl.SetPackage)

	// * Healthcheck
	api.Get("/healthcheck", ctrl.Healthcheck)
	api.Get("/_healthcheck", ctrl.InternalHealthcheck)

	// * Authentication group
	auth := api.Group("auth")
	// auth.Get("/publickey", authCtrl.PublicKey)
	auth.Get("/verify", authCtrl.Verify)
	auth.Post("/login", authCtrl.Login)
	auth.Post("/register", authCtrl.Register)
	auth.Post("/refresh", mid.AuthMiddleware, authCtrl.Refresh)
	auth.Post("/logout", mid.AuthMiddleware, authCtrl.Logout)
	auth.Post("/genreset", authCtrl.GenerateReset)
	auth.Get("/reset", authCtrl.GetReset)
	auth.Post("/reset", authCtrl.PostReset)

	// * MINIO S3
	auth.Post("/minio", authCtrl.Minio)
	auth.Head("/minio", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusNoContent) })

	// * User group
	user := api.Group("users", mid.AuthMiddleware, mid.AdminMiddleware)
	user.Get("/:id", usersCtrl.GetUser)
	user.Get("/", usersCtrl.GetUsers)
	user.Delete("/:id", usersCtrl.DeleteUser)
	user.Patch("/:id", usersCtrl.PatchUser)
	// User packages
	packages := user.Group("packages")
	packages.Get("/:id", packagesCtrl.GetPackage)
	packages.Delete("/pack/:id", packagesCtrl.DeletePackage)
	packages.Patch("/pack/:id", packagesCtrl.PatchPackage)
	packages.Post("/", packagesCtrl.PostPackage)

	// * Profile
	profile := api.Group("profile", mid.AuthMiddleware)
	profile.Get("/package", profileCtrl.GetPackage)
	profile.Post("/pay", profileCtrl.PostPayment)

	// * Video group
	video := api.Group("video", mid.AuthMiddleware)
	video.Post("/", mid.AdminMiddleware, videoCtrl.PostVideo)
	video.Get("/all", mid.AdminMiddleware, videoCtrl.GetAllVideos)
	video.Get("/:id", videoCtrl.GetVideo)
	video.Get("/", videoCtrl.GetRecommendedVideos)
	video.Delete("/:id", videoCtrl.DeleteVideo)
	video.Patch("/:id", videoCtrl.PathVideo)

	// * Category group
	category := video.Group("category")
	category.Get("/All", videoCategoryCtrl.GetCategories)
	category.Get("/:id", videoCategoryCtrl.GetCategory)
	category.Patch("/:id", mid.AdminMiddleware, videoCategoryCtrl.PatchCategory)
	category.Post("/", mid.AdminMiddleware, videoCategoryCtrl.PostCategory)
	category.Delete("/:id", mid.AdminMiddleware, videoCategoryCtrl.DeleteCategory)

	// * Administration
	// admin := api.Group("admin", mid.AuthMiddleware)

	log.Fatal(app.Listen(":3000"))
}

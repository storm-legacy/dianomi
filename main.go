package main

import (
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/storm-legacy/dianomi/internal/controllers"
	"github.com/storm-legacy/dianomi/internal/database"

	log "github.com/sirupsen/logrus"
)

// Utility functions
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func init() {
	// Specify log level
	log.SetLevel(log.DebugLevel)

	// Load .env if exists to
	err := godotenv.Load(".env")
	if err != nil {
		log.WithField("value", err).Warn("Could not load environment file.")
	}

	// Connect to database & migrate
	database.ConnectToDatabase(os.Getenv("DATABASE_URL"))
	database.AutoMigrate()

}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Healthcheck
	app.Get("/api/v1/healthcheck", func(c *fiber.Ctx) error {
		// if c.IP() == "127.0.0.1" {
		if true {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"healthy": true,
				"resources": fiber.Map{
					"allocMb":      bToMb(m.Alloc),
					"totalAllocMb": bToMb(m.TotalAlloc),
					"sysMb":        bToMb(m.Sys),
					"numGCMb":      bToMb(uint64(m.NumGC)),
				},
			})
		}
		return c.SendStatus(fiber.StatusForbidden)
	})

	app.Post("/api/v1/auth/register", controllers.RegisterUser)
	app.Post("/api/v1/auth/login", controllers.LoginUser)

	log.Fatal(app.Listen(":3000"))
}

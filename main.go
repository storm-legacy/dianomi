package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/storm-legacy/dianomi/internal/controllers"
	"github.com/storm-legacy/dianomi/internal/database"

	log "github.com/sirupsen/logrus"
)

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

	app.Get("/api/v1/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"healthy": true,
		})
	})

	app.Post("/api/v1/auth/register", controllers.RegisterUser)

	log.Fatal(app.Listen(":3000"))
}

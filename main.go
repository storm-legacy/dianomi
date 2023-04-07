package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
	"github.com/storm-legacy/dianomi/internal/database"

	log "github.com/sirupsen/logrus"
)

var (
	privateKey ed25519.PrivateKey

	publicKey       ed25519.PublicKey
	publicKeyBase64 string
)

func init() {
	// Specify log level
	log.SetLevel(log.DebugLevel)

	// Load configuration from file
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.WithField("msg", err.Error()).Info("Configuration file .env doesn't exists")
	}

	// Check for ed25519 private key and generate if needed
	// TODO Add error handler
	pkey := viper.GetString("APP_JWT_EDDSA_PRIVATE_KEY")
	privateKey, _ = base64.StdEncoding.DecodeString(pkey)
	publicKeyBase64 = viper.GetString("APP_JWT_EDDSA_PUBLIC_KEY")
	publicKey, _ = base64.StdEncoding.DecodeString(publicKeyBase64)

	// Generate key pair if any empty
	if len(privateKey) != 64 || len(publicKey) != 32 {
		publicKey, privateKey, _ = ed25519.GenerateKey(rand.Reader)
		publicKeyBase64 = base64.StdEncoding.EncodeToString(publicKey)
		log.WithField("publicKey", publicKeyBase64).Info("Generated ed25519 key pair for authentication")
	}

	// Connect to database & migrate
	databaseURL := fmt.Sprintf(
		`postgresql://%s:%s@%s:%d/%s`,
		viper.Get("APP_PG_USER"),
		viper.Get("APP_PG_PASSWORD"),
		viper.Get("APP_PG_HOST"),
		viper.GetInt("APP_PG_PORT"),
		viper.Get("APP_PG_DB"),
	)
	database.ConnectToDatabase(databaseURL)
	database.Migrate()

}

func main() {
	app := fiber.New()
	app.Use(logger.New(), func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Get("/api/v1/auth/publickey", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"publicKey": publicKeyBase64,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

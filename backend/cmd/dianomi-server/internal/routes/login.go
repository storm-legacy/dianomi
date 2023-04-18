package routes

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/cmd/dianomi-server/internal/handlers"
)

func Login(c *fiber.Ctx) error {

	// Parse user data
	var userData *handlers.FormLoginUser
	if err := c.BodyParser(&userData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Handle login
	result := handlers.LoginUser(userData)

	switch result.Status() {

	// Operation was success
	case handlers.SUCCESS:
		log.WithField("email", userData.Email).Debug("user's token generated")
		return c.Status(result.StatusCode).JSON(result.Data)

	// Client made error
	case handlers.CLIENT_ERROR:
		log.WithField("message", result.ErrorMessage).Debug("client error occured")
		return c.Status(result.StatusCode).JSON(fiber.Map{"error": result.ErrorMessage})

	// Server error
	case handlers.SERVER_ERROR:
		log.WithField("error", result.ErrorMessage).Error("internal server error occured")
		return c.Status(result.StatusCode).JSON(fiber.Map{"error": "internal server error"})
	}

	// Occured error without handler for it
	log.WithFields(log.Fields{"error": result.ErrorMessage}).Error("unknown error occured")
	return c.SendStatus(result.StatusCode)
}

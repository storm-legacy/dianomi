package controllers

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/internal/argon2"
	db "github.com/storm-legacy/dianomi/internal/database"
	"github.com/storm-legacy/dianomi/internal/models"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func RegisterUser(c *fiber.Ctx) error {
	// Model for storing data
	var newUserData *models.RegisterUser

	// Parse data and return problem if occured
	if err := c.BodyParser(&newUserData); err != nil {
		log.WithFields(log.Fields{
			"func": "RegisterUser()",
			"err":  err,
		}).Debug("Wrong model used for registration")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err,
		})
	}

	// Check if email is correct
	if !emailRegex.MatchString(newUserData.Email) {
		log.WithFields(log.Fields{
			"func":  "RegisterUser()",
			"email": newUserData.Email,
		}).Debug("Incorrect email used for registration")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Incorrect email address",
		})
	}

	// Check if email already exists
	result, err := db.IsNewUser(newUserData)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "RegisterUser()",
			"err":  err,
		}).Error("Problem while checking for user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal error",
		})
	} else if !result {
		log.WithFields(log.Fields{
			"func":  "RegisterUser()",
			"email": newUserData.Email,
		}).Debug("User already exists")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
		})
	}

	// Check if password checks minimal length
	if len(newUserData.Password) < 8 {
		log.WithFields(log.Fields{
			"func":  "RegisterUser()",
			"email": newUserData.Email,
		}).Debug("Password is too simple")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is too simple",
		})
	}

	// Check if password is the same as repeated one
	if newUserData.Password != newUserData.PasswordRepeated {
		log.WithFields(log.Fields{
			"func":  "RegisterUser()",
			"email": newUserData.Email,
		}).Debug("Passwords are not the same")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Passwords are not the same",
		})
	}

	// hash password
	hashedPassword, err := argon2.EncodePassword(&newUserData.Password)
	if err != nil {
		log.WithFields(log.Fields{
			"func": "RegisterUser()",
			"err":  err,
		}).Error("Problem occured while hashing password")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal error",
		})
	}

	// Create user
	user := models.User{
		Email:    newUserData.Email,
		Password: hashedPassword,
	}

	// Insert to database
	if err = db.CreateUser(&user); err != nil {
		log.WithFields(log.Fields{
			"func": "RegisterUser()",
			"err":  err,
		}).Error("Could not add user to database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Internal error",
		})
	}

	// Success
	log.WithField("email", newUserData.Email).Debug("New user registered")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"email": newUserData.Email,
		}},
	)
}

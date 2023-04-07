package middlewares

import (
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	db "github.com/storm-legacy/dianomi/internal/database"
	"github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/argon2"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

var (
	jwtExpiredIn string = "60m"
	jwtMaxAge    int    = 60
	jwtSecret    string = "changeme123"
)

// FIXME Adjust log levels
func RegisterUser(c *fiber.Ctx) error {
	// Model for storing data
	var registerUser *models.RegisterUser

	// Parse data and return problem if occured
	if err := c.BodyParser(&registerUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})
	}

	// Change email to lowercase
	registerUser.Email = strings.ToLower(registerUser.Email)

	// Check if email is correct
	if !emailRegex.MatchString(registerUser.Email) {
		log.WithFields(log.Fields{"func": "RegisterUser()", "email": registerUser.Email}).Debug("Incorrect email used for registration")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Incorrect email address",
		})
	}

	// Check if email already exists
	result, err := db.IsNewUser(registerUser)
	if err != nil {
		log.WithFields(log.Fields{"func": "RegisterUser()", "err": err}).Error("Problem while checking for user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})

	} else if !result {
		log.WithFields(log.Fields{"func": "RegisterUser()", "email": registerUser.Email}).Debug("User already exists")
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status":  "error",
			"message": "User already exists",
		})
	}

	// Check if password checks minimal length
	if len(registerUser.Password) < 8 {
		log.WithFields(log.Fields{"func": "RegisterUser()", "email": registerUser.Email}).Debug("Password is too simple")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Password is too simple",
		})
	}

	// Check if password is the same as repeated one
	if registerUser.Password != registerUser.PasswordRepeated {
		log.WithFields(log.Fields{"func": "RegisterUser()", "email": registerUser.Email}).Debug("Passwords are not the same")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Passwords are not the same",
		})
	}

	// hash password
	hashedPassword, err := argon2.EncodePassword(&registerUser.Password)
	if err != nil {
		log.WithFields(log.Fields{"func": "RegisterUser()", "err": err}).Error("Problem occured while hashing password")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})
	}

	// Create user
	user := models.User{
		Email:    registerUser.Email,
		Password: hashedPassword,
	}

	// Insert to database
	if err = db.CreateUser(&user); err != nil {
		log.WithFields(log.Fields{"func": "RegisterUser()", "err": err}).Error("Could not add user to database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})
	}

	// Success
	log.WithField("email", registerUser.Email).Debug("New user registered")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"email": registerUser.Email,
		}},
	)
}

func LoginUser(c *fiber.Ctx) error {
	var loginUser *models.LoginUser

	// Parse data and return problem if occured
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})
	}

	// Check if email and password are valid
	if !emailRegex.MatchString(loginUser.Email) || len(loginUser.Password) < 8 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Incorrect credentials"})
	}

	// Get user by email from database
	user, err := db.GetUser(loginUser.Email)
	if err != nil {
		log.WithFields(log.Fields{"func": "LoginUser()", "err": err}).Error("Error occured while getting the user from database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})

		// User doesn't exists
	} else if user == nil {
		log.WithFields(log.Fields{"func": "LoginUser()", "email": user.Email}).Debug("User does not exists")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Incorrect credentials"})
	}

	// Compare password of the user
	result, err := argon2.ComparePasswordAndHash(&loginUser.Password, &user.Password)
	if err != nil {
		log.WithFields(log.Fields{"func": "LoginUser()", "err": err}).Error("Problem while comparing passwords")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})

		// Password is incorrect
	} else if !result {
		log.WithFields(log.Fields{"func": "LoginUser()", "email": loginUser.Email}).Debug("Used incorrect password")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Incorrect credentials"})
	}

	duration, _ := time.ParseDuration(jwtExpiredIn)

	token := jwt.New(jwt.SigningMethodHS256)
	now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(duration).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.WithFields(log.Fields{"func": "LoginUser()", "err": err}).Error("Could not generate JWT token")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal error"})
	}

	c.Cookie(&fiber.Cookie{
		Name:   "token",
		Value:  tokenString,
		Path:   "/",
		MaxAge: jwtMaxAge * 60,
		Secure: true,
	})

	// Success
	log.WithField("email", user.Email).Debug("User logged in")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  tokenString,
	})
}

package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type passwordData struct {
	Email       string `json:"email" validate:"required"`
	OldPassword string `json:"OldPassword" validate:"required"`
	NewPassword string `json:"NewPassword"`
}
type userEmail struct {
	Email       string `json:"email" validate:"required"`
	OldPassword string `json:"OldPassword" validate:"required"`
}

func GetOldPassword(c *fiber.Ctx) error {
	var email userEmail
	if err := c.BodyParser(&email); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	queries := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)
	user, err := queries.GetUserByEmail(ctx, email.Email)
	if err == sql.ErrNoRows {
		log.WithField("user", email.Email).Debug("User doesn't exist")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
		})
	}
	if err != nil {
		log.WithField("err", err).Error("SQL query resulted in error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	result, err := comparePasswords(email.OldPassword, user.Password)
	if err != nil {
		log.WithField("err", err).Error("Problem decoding password")
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if !result {
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
		})
	}
	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data: fiber.Map{
			"result": result,
		},
	})
}

func comparePasswords(password string, hashedPassword string) (bool, error) {
	result, err := argon2.ComparePasswordAndHash(&password, &hashedPassword)
	if err != nil {
		return false, err
	}
	return result, nil
}
func PatchPasswortResetProfil(c *fiber.Ctx) error {

	var data passwordData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	hashedPassword, err := argon2.EncodePassword(&data.OldPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	result, err := comparePasswords(data.OldPassword, hashedPassword)
	if err != nil {
		log.WithField("err", err).Error("Problem decoding password")
		return c.SendStatus(fiber.StatusInternalServerError)
	} else if !result {
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
		})
	}

	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data: fiber.Map{
			"Email":       hashedPassword,
			"OldPassword": data.OldPassword,
			"NewPassword": result,
		},
	})
}

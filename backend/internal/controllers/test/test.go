package controllers

import (
	"context"
	"database/sql"
	"fmt"

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
			Data:   "Incorrect information",
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

func NewUserPassword(c *fiber.Ctx) error {

	var data passwordData
	if err := c.BodyParser(&data); err != nil {
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

	user, err := queries.GetUserByEmail(ctx, data.Email)
	if err == sql.ErrNoRows {
		log.WithField("user", data.Email).Debug("User doesn't exist")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect information",
		})
	}

	hashedPassword, err := argon2.EncodePassword(&data.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	fmt.Println(data.NewPassword)

	err = queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:       user.ID,
		Password: hashedPassword,
	})
	fmt.Println(comparePasswords(data.NewPassword, hashedPassword))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)

}

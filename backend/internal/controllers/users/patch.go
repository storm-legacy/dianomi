package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	authCtrl "github.com/storm-legacy/dianomi/internal/controllers/auth"
	"github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type PatchUserData struct {
	Email         string `json:"email" validate:"required,email"`
	Verified      bool   `json:"verified" validate:"boolean"`
	ResetPassword bool   `json:"reset_password" validate:"boolean"`
}

func PatchUser(c *fiber.Ctx) error {
	// Get video ID
	idString := c.Params("id")
	paramId, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Parse sent data
	var data PatchUserData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	// Validate data
	err = models.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	// * CHECK AGANIST DATABASE
	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	qtx := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	// Check if exists
	user, err := qtx.GetUserByID(ctx, paramId)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("userID", user.ID).Debug("User with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Email update
	if err := qtx.UpdateUserEmail(ctx, sqlc.UpdateUserEmailParams{
		ID:    user.ID,
		Email: data.Email,
	}); err != nil {
		log.WithField("userID", user.ID).Error("Users email could not be updated")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Verification update
	if data.Verified && !user.VerifiedAt.Valid {
		now := sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		if err := qtx.UpdateUserVerification(ctx, sqlc.UpdateUserVerificationParams{
			ID:         user.ID,
			VerifiedAt: now,
		}); err != nil {
			log.WithField("userID", user.ID).Error("User could not be verified")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	} else if !data.Verified && user.VerifiedAt.Valid {
		nullvalue := sql.NullTime{
			Time:  time.Now(),
			Valid: false,
		}
		if err := qtx.UpdateUserVerification(ctx, sqlc.UpdateUserVerificationParams{
			ID:         user.ID,
			VerifiedAt: nullvalue,
		}); err != nil {
			log.WithField("userID", user.ID).Error("User could not be verified")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// Trigger password reset
	if data.ResetPassword {
		go authCtrl.AsyncReset(authCtrl.ResetData{
			Email: user.Email,
		})
	}

	return c.SendStatus(fiber.StatusOK)
}

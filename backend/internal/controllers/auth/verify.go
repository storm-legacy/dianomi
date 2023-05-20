package controllers

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func Verify(c *fiber.Ctx) error {
	frontUrl := config.GetString("APP_FRONT_URL")
	queryUuid := c.Query("validate")

	// * CHECK AGANIST DATABASE
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

	// Check if uuid is correct
	uuidValue, err := uuid.Parse(queryUuid)
	if err != nil {
		log.WithField("err", err).Debug("Send UUID is not correct")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if is valid
	uuidDb, err := queries.GetVerificationCode(ctx, uuidValue)
	if err == sql.ErrNoRows {
		log.WithField("uuid", uuidDb).Debug("Provided uuid doesn't exist in database")
		return c.SendStatus(fiber.StatusBadRequest)
	} else if err != nil {
		log.WithField("err", err).Error("Could not get verification code from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Check if wasn't used
	now := time.Now().Unix()
	if uuidDb.ValidUntil.Time.Unix() < now {

		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Verify user
	if err := queries.VerifyUser(ctx, uuidDb.UserID); err != nil {
		log.WithField("err", err).Error("Could not verify the user")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Invalidate code for later requests
	if err := queries.SetCodeAsUsed(ctx, uuidDb.ID); err != nil {
		log.WithField("err", err).Error("Code could not be set as used")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).Redirect(frontUrl)
}

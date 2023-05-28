package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func PatchPackage(c *fiber.Ctx) error {
	var data Package
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	// Validate data
	err := models.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}
	if data.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing package id",
		})
	}

	// Check if package is correct
	if data.ValidFrom.Unix() > data.ValidUntil.Unix() {
		log.WithFields(log.Fields{
			"validFrom":  data.ValidFrom,
			"validUntil": data.ValidUntil,
		}).Warn("Time in one or more packages is incorrect")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ValidFrom and ValidUntil values aren't correct",
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

	// Update package
	if err := qtx.UpdatePackage(ctx, sqlc.UpdatePackageParams{
		ID:         data.ID,
		UserID:     sql.NullInt64{Int64: int64(data.UserID), Valid: true},
		Tier:       sqlc.Role(data.Tier),
		ValidFrom:  data.ValidFrom,
		ValidUntil: data.ValidUntil,
	}); err != nil {
		log.WithField("err", err.Error()).Error("Could not give user the package")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

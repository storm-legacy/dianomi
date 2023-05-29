package controllers

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type Package struct {
	ID         int64     `json:"id"`
	UserID     uint64    `json:"user_id" validate:"required"`
	Tier       string    `json:"tier" validate:"required"`
	ValidFrom  time.Time `json:"valid_from" validate:"required"`
	ValidUntil time.Time `json:"valid_until" validate:"required"`
}

func GetPackage(c *fiber.Ctx) error {
	id := c.Locals("sub").(uint64)

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

	// Get user packages
	packs, err := qtx.GetPackagesByUserID(ctx, sqlc.GetPackagesByUserIDParams{
		UserID: sql.NullInt64{Int64: int64(id), Valid: true},
		Limit:  1,
		Offset: 0,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get user packages from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if len(packs) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": Package{
			ID:         packs[0].ID,
			UserID:     uint64(packs[0].UserID.Int64),
			Tier:       string(packs[0].Tier),
			ValidFrom:  packs[0].ValidFrom,
			ValidUntil: packs[0].ValidUntil,
		},
	})
}

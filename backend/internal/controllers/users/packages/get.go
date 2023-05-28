package controllers

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func GetPackage(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	limitString := c.Query("limit")
	limit, err := strconv.ParseInt(string(limitString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("Limit could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	pageString := c.Query("page")
	page, err := strconv.ParseInt(string(pageString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("Page could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if page > 0 {
		page -= 1
	}

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

	packs, err := qtx.GetPackagesByUserID(ctx, sqlc.GetPackagesByUserIDParams{
		UserID: sql.NullInt64{Int64: id, Valid: true},
		Limit:  int32(limit),
		Offset: int32(page * limit),
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get user packages from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	packages := make([]Package, 0)
	for _, p := range packs {
		var newPack = Package{
			ID:         p.ID,
			UserID:     uint64(p.UserID.Int64),
			Tier:       string(p.Tier),
			ValidFrom:  p.ValidFrom,
			ValidUntil: p.ValidUntil,
		}
		packages = append(packages, newPack)
	}

	return c.Status(fiber.StatusOK).JSON(packages)
}

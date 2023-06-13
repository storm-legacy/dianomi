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

func PostBanUser(c *fiber.Ctx) error {
	// Get video ID
	idString := c.Params("id")
	paramId, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
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

	// Ban user
	if err := qtx.BanUser(ctx, paramId); err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("User couldn't be banned")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func PostUnbanUser(c *fiber.Ctx) error {
	// Get video ID
	idString := c.Params("id")
	paramId, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
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

	// Ban user
	if err := qtx.UnbanUser(ctx, paramId); err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("User couldn't be banned")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

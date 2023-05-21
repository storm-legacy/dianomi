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

func DeleteUser(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
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

	user, err := qtx.GetUserByID(ctx, id)
	if err == sql.ErrNoRows {
		log.WithField("user", id).Debug("Specified user doesn't exist")
		return c.SendStatus(fiber.StatusNotModified)
	}
	if err != nil {
		log.WithField("err", err).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Delete from database
	err = qtx.DeleteUser(ctx, user.ID)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error(), "user_id": id}).Debug("User could not be removed from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

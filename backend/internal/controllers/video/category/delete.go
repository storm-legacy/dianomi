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

func DeleteCategory(c *fiber.Ctx) error {

	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
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

	// Check if exists
	_, err = qtx.GetCategoryByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get categories from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("category_id", id).Debug("Category with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotModified)
	}

	// Delete from database
	err = qtx.DeleteCategory(ctx, id)
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error(), "category_id": id}).Debug("Category could not be removed")
		return c.SendStatus(fiber.StatusNotModified)
	}

	return c.SendStatus(fiber.StatusOK)
}

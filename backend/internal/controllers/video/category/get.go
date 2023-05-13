package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func GetCategory(c *fiber.Ctx) error {

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
	category, err := qtx.GetCategoryByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get categories from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("category_id", id).Debug("Category with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

func GetCategories(c *fiber.Ctx) error {
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

	// Check if limit and offset is resonable
	var offset int32 = 0
	offsetArray := c.Query("offset")
	fmt.Print(string(offsetArray))
	if len(offsetArray) > 1 {
		result, err := strconv.ParseInt(string(offsetArray), 10, 64)
		if err != nil {
			log.WithField("err", err).Debug("Offset value could not be parsed")
			return c.SendStatus(fiber.StatusBadRequest)
		}
		offset = int32(result)
	}

	res, err := qtx.GetAllCategories(ctx, sqlc.GetAllCategoriesParams{
		Limit:  25,
		Offset: offset,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get categories from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

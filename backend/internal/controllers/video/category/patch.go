package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type CategoryPatchData struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func PatchCategory(c *fiber.Ctx) error {

	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var patchedData CategoryPatchData
	patchedData.Id = -1
	if err := c.BodyParser(&patchedData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = mod.Validate.Struct(patchedData)
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
	category, err := qtx.GetCategoryByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get categories from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("category_id", id).Debug("Category with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Compare values
	if category.Name == patchedData.Name {
		return c.SendStatus(fiber.StatusNotModified)
	}

	if err = qtx.UpdateCategory(ctx, sqlc.UpdateCategoryParams{
		ID:   category.ID,
		Name: patchedData.Name,
	}); err != nil {
		log.WithField("category_id", id).Error("Problem with updating category")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

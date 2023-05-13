package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type CategoryPostData struct {
	Name string `json:"name" validate:"required,max=32"`
}

func PostCategory(c *fiber.Ctx) error {
	var categoryData CategoryPostData
	if err := c.BodyParser(&categoryData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := mod.Validate.Struct(categoryData)
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
	_, err = qtx.GetCategoryByName(ctx, categoryData.Name)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get categories from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err != sql.ErrNoRows {
		log.WithFields(log.Fields{"category": categoryData.Name}).Debug("Category already exists")
		return c.SendStatus(fiber.StatusConflict)
	}

	// Add category
	_, err = qtx.AddCategory(ctx, categoryData.Name)
	if err != nil {
		log.WithField("err", err.Error()).Error("Category could not be added")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

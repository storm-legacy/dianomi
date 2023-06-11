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

type PostCommentsData struct {
	Email   string `json:"email" validate:"required"`
	VideoID int64  `json:"video_id" validate:"required"`
	Comment string `json:"comment" validate:"required" `
}

func PostComments(c *fiber.Ctx) error {
	var data PostCommentsData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
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

	user, err := qtx.GetUserByEmail(ctx, data.Email)
	if err == sql.ErrNoRows {
		log.WithField("user", data.Email).Debug("User doesn't exist")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect information",
		})
	}
	if err != nil {
		log.WithField("err", err).Error("SQL query resulted in error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := qtx.AddComments(ctx, sqlc.AddCommentsParams{
		UserID:  user.ID,
		VideoID: data.VideoID,
		Comment: data.Comment,
	}); err != nil {
		log.WithField("err", err.Error()).Error("Comment could not be added")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}

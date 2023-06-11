package controllers

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type GetCommentsForVideoData struct {
	Email     string    `json:"email" validate:"required"`
	Comment   string    `json:"comment" validate:"required" `
	UpdatedAt time.Time `json:"updated_at"`
}

func GetCommentsVideo(c *fiber.Ctx) error {
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

	res, err := qtx.GetCommentsForVideo(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get comments from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("video_id", id).Debug("Specified comments doesn't exist")
		return c.SendStatus(fiber.StatusNotModified)
	}
	commentsData := make([]GetCommentsForVideoData, 0)
	for _, met := range res {
		comment := GetCommentsForVideoData{
			Email:     met.Email,
			Comment:   met.Comment,
			UpdatedAt: met.UpdatedAt.Time,
		}
		commentsData = append(commentsData, comment)
	}

	return c.Status(fiber.StatusOK).JSON(commentsData)

}

package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type VideoPatchData struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	CategoryId  int64    `json:"category_id" validate:"required"`
	Premium     bool     `json:"is_premium"`
	Tags        []string `json:"tags" validate:"required,tags"`
}

func PathVideo(c *fiber.Ctx) error {
	// Parse sent data
	var data VideoPatchData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// Get video ID
	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = mod.Validate.Struct(data)
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
	video, err := qtx.GetVideoByID(ctx, id)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get video from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("video_id", id).Debug("Video with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Update video
	if _, err := qtx.UpdateVideo(ctx, sqlc.UpdateVideoParams{
		ID:          video.ID,
		Name:        data.Name,
		Description: data.Description,
		CategoryID:  data.CategoryId,
		IsPremium:   data.Premium,
	}); err != nil {
		log.WithField("video_id", id).Error("Problem with updating the video")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Remove tags from video
	if err := qtx.ClearVideoTags(ctx, video.ID); err != nil {
		log.WithField("video_id", id).Error("Tags could not be removed from video")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Update tags
	// Trim tags and check if exists
	// Add to database if needed and get ids to add to video
	var tagIds []int64
	for _, tag := range data.Tags {
		tag = strings.TrimSpace(tag)
		tag = strings.ToLower(tag)

		dbTag, err := qtx.GetTagByName(ctx, tag)
		if err == sql.ErrNoRows {
			dbTag, err = qtx.AddTag(ctx, tag)
			if err != nil {
				log.WithField("err", err.Error()).Error("Could not add tag to database")
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		} else if err != nil {
			log.WithField("err", err.Error()).Error("Problem getting tags from database")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		tagIds = append(tagIds, dbTag.ID)
	}

	// Attach tags to video
	for _, tagId := range tagIds {
		if err := qtx.AddVideoTag(ctx, sqlc.AddVideoTagParams{
			VideoID: video.ID,
			TagID:   tagId,
		}); err != nil {
			log.WithField("err", err.Error()).Error("Problem with attaching tag to video")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

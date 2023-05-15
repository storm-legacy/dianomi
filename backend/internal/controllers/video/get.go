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

type VideoResponse struct {
	ID           uint64   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Category     string   `json:"category"`
	CategoryID   uint64   `json:"category_id"`
	Upvotes      uint64   `json:"upvotes"`
	Downvotes    uint64   `json:"downvotes"`
	Views        uint64   `json:"views"`
	ThumbnailUrl string   `json:"thumbnail_url"`
	Tags         []string `json:"tags"`
}

func GetAllVideos(c *fiber.Ctx) error {
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

	res, err := qtx.GetAllVideos(ctx, sqlc.GetAllVideosParams{
		Limit:  25,
		Offset: 0,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get videos from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	videos := make([]VideoResponse, 0)
	for _, vid := range res {
		tags, err := qtx.GetVideoTags(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		video := VideoResponse{
			ID:           uint64(vid.ID),
			Name:         vid.Name,
			Description:  vid.Description,
			Category:     vid.Category.String,
			CategoryID:   uint64(vid.ID),
			Upvotes:      uint64(vid.Upvotes),
			Downvotes:    uint64(vid.Downvotes),
			Views:        uint64(vid.Views),
			ThumbnailUrl: vid.Thumbnail.String,
			Tags:         tags,
		}
		videos = append(videos, video)
	}

	return c.Status(fiber.StatusOK).JSON(videos)
}

func GetRecommendedVideos(c *fiber.Ctx) error {
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

	res, err := qtx.GetRandomVideos(ctx, sqlc.GetRandomVideosParams{
		Limit:  25,
		Offset: offset,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get videos from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	videos := make([]VideoResponse, 0)
	for _, vid := range res {
		tags, err := qtx.GetVideoTags(ctx, vid.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get tags from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		video := VideoResponse{
			ID:           uint64(vid.ID),
			Name:         vid.Name,
			Description:  vid.Description,
			Category:     vid.Category.String,
			CategoryID:   uint64(vid.ID),
			Upvotes:      uint64(vid.Upvotes),
			Downvotes:    uint64(vid.Downvotes),
			Views:        uint64(vid.Views),
			ThumbnailUrl: vid.Thumbnail.String,
			Tags:         tags,
		}
		videos = append(videos, video)
	}

	return c.Status(fiber.StatusOK).JSON(videos)
}

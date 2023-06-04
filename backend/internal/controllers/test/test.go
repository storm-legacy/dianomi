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

type Email struct {
	Email string `json:"email" validate:"required"`
}

type UserVideoMetricData struct {
	ID                int64        `json:"id"`
	UserID            int64        `json:"user_id"`
	VideoID           int64        `json:"video_id"`
	TimeSpentWatching int64        `json:"time_spent_watching"`
	StoppedAt         int64        `json:"stopped_at"`
	CreatedAt         sql.NullTime `json:"created_at"`
	UpdatedAt         sql.NullTime `json:"updated_at"`
}

type PostVideoMerticsData struct {
	Email             string `json:"email" validate:"required"`
	VideoID           int64  `json:"video_id" validate:"required"`
	TimeSpentWatching int64  `json:"time_spent_watching" validate:"required"`
	StoppedAt         int64  `json:"stopped_at" validate:"required"`
}

func PostVideoMertics(c *fiber.Ctx) error {

	var data PostVideoMerticsData
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}
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

	userMetric, err := qtx.IfUserSeeThisVideo(ctx, sqlc.IfUserSeeThisVideoParams{
		UserID:  user.ID,
		VideoID: data.VideoID,
	})

	if userMetric == (sqlc.UserVideoMetric{}) {
		if err := qtx.AddVideoMertics(ctx, sqlc.AddVideoMerticsParams{
			UserID:            user.ID,
			VideoID:           data.VideoID,
			TimeSpentWatching: int32(data.TimeSpentWatching),
			StoppedAt:         int32(data.StoppedAt),
		}); err != nil {
			log.WithField("err", err.Error()).Error("Video couldn't be added to history")
			return c.SendStatus(fiber.StatusAccepted)

		}

		return c.SendStatus(fiber.StatusAccepted)
	} else {
		if err := qtx.UpdateVideoMetric(ctx, sqlc.UpdateVideoMetricParams{
			TimeSpentWatching: int32(data.TimeSpentWatching),
			StoppedAt:         int32(data.StoppedAt),
			ID:                userMetric.ID,
		}); err != nil {
			log.WithField("err", err.Error()).Error("Video couldn't be added to history")
			return c.SendStatus(fiber.StatusAccepted)

		}
		return c.SendStatus(fiber.StatusAccepted)

	}

}

func GetUserVideoMertics(c *fiber.Ctx) error {
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

	res, err := qtx.GetAllVideoMetric(ctx, sqlc.GetAllVideoMetricParams{
		Limit:  25,
		Offset: 0,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get video metric from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	videoMetrics := make([]UserVideoMetricData, 0)
	for _, met := range res {

		videoMetric := UserVideoMetricData{
			ID:                int64(met.ID),
			UserID:            int64(met.UserID),
			VideoID:           int64(met.VideoID),
			TimeSpentWatching: int64(met.TimeSpentWatching),
			StoppedAt:         int64(met.StoppedAt),
			CreatedAt:         met.CreatedAt,
			UpdatedAt:         met.UpdatedAt,
		}

		videoMetrics = append(videoMetrics, videoMetric)
	}
	return c.Status(fiber.StatusOK).JSON(videoMetrics)

}

func GetUserVideoMerticByEmail(c *fiber.Ctx) error {
	var data Email
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}
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
	res, err := qtx.GetUserVideoMerticsByUserId(ctx, user.ID)
	if err != nil {
		log.WithField("err", err).Error("Could not get video metric from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	videoMetrics := make([]UserVideoMetricData, 0)
	for _, met := range res {

		videoMetric := UserVideoMetricData{
			ID:                int64(met.ID),
			UserID:            int64(met.UserID),
			VideoID:           int64(met.VideoID),
			TimeSpentWatching: int64(met.TimeSpentWatching),
			StoppedAt:         int64(met.StoppedAt),
			CreatedAt:         met.CreatedAt,
			UpdatedAt:         met.UpdatedAt,
		}

		videoMetrics = append(videoMetrics, videoMetric)
	}
	return c.Status(fiber.StatusOK).JSON(videoMetrics)
}

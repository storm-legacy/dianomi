package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type testData struct {
	Test   string `json:"test"`
	Number int32  `json:"Numer"`
}

func Test(c *fiber.Ctx) error {
	var dataTest *testData
	if err := c.BodyParser(&dataTest); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if dataTest.Number == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "error",
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

	user, err := qtx.GetUserByID(ctx, int64(dataTest.Number))
	if err != nil {
		log.WithField("err", err).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

package controllers

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type DevPackageInput struct {
	Email    string    `json:"email"`
	Role     sqlc.Role `json:"role"`
	ValidFor int32     `json:"valid_for"`
}

func SetPackage(c *fiber.Ctx) error {
	// * PARSE DATA
	var postData DevPackageInput
	if err := c.BodyParser(&postData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// * VALIDATE DATA
	postData.Email = strings.ToLower(postData.Email)
	postData.Email = strings.TrimSpace(postData.Email)
	// validate data
	err := mod.Validate.Struct(postData)
	if err != nil {
		log.WithField("err", err).Debug("Could not validate user data")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect registration information provided",
		})
	}

	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	queries := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	user, err := queries.GetUserByEmail(ctx, postData.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err = queries.RemoveAllUserPackages(ctx, user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	err = queries.GiveUserPackage(ctx, sqlc.GiveUserPackageParams{
		UserID:     user.ID,
		Tier:       postData.Role,
		ValidFrom:  time.Now(),
		ValidUntil: time.Now().AddDate(0, 0, int(postData.ValidFor)),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

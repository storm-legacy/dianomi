package controllers

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type User struct {
	ID       uint64 `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Verified bool   `json:"verified" validate:"required"`
	Role     string `json:"role"`
}

func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

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

	user, err := qtx.GetUserByEmail(ctx, email)
	if err != nil {
		log.WithField("err", err).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err == sql.ErrNoRows {
		log.WithField("email", email).Debug("Specified user doesn't exist")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	pkg, err := qtx.GetCurrentPackageForUser(ctx, sql.NullInt64{Int64: user.ID, Valid: true})
	if err != nil && err != sql.ErrNoRows {
		log.WithField("err", err).Error("Could not get user packages from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	role := "free"
	if err != sql.ErrNoRows {
		role = string(pkg.Tier)
	}

	return c.Status(fiber.StatusOK).JSON(User{
		ID:       uint64(user.ID),
		Email:    user.Email,
		Verified: user.VerifiedAt.Valid,
		Role:     role,
	})
}

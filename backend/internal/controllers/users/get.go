package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type Package struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"user_id"`
	Tier       string    `json:"tier"`
	ValidFrom  time.Time `json:"valid_from"`
	ValidUntil time.Time `json:"valid_until"`
}

type User struct {
	ID       uint64    `json:"id"`
	Email    string    `json:"email"`
	Verified bool      `json:"verified"`
	Packages []Package `json:"packages"`
}

func GetUser(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
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

	user, err := qtx.GetUserByID(ctx, id)
	if err != nil {
		log.WithField("err", err).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err == sql.ErrNoRows {
		log.WithField("user", id).Debug("Specified user doesn't exist")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	packs, err := qtx.GetPackagesByUserID(ctx, user.ID)
	if err != nil {
		log.WithField("err", err).Error("Could not get user packages from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}
	packages := make([]Package, 0)
	for _, p := range packs {
		var newPack = Package{
			ID:         uint64(p.ID),
			Tier:       string(p.Tier),
			ValidFrom:  p.ValidFrom,
			ValidUntil: p.ValidUntil,
		}
		packages = append(packages, newPack)
	}

	return c.Status(fiber.StatusOK).JSON(User{
		ID:       uint64(user.ID),
		Email:    user.Email,
		Verified: user.VerifiedAt.Valid,
		Packages: packages,
	})
}

func GetUsers(c *fiber.Ctx) error {
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

	res, err := qtx.GetAllUsers(ctx, sqlc.GetAllUsersParams{
		Limit:  25,
		Offset: offset,
	})
	if err != nil {
		log.WithField("err", err).Error("Could not get users from database")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	users := make([]User, 0)
	for _, user := range res {
		packs, err := qtx.GetPackagesByUserID(ctx, user.ID)
		if err != nil {
			log.WithField("err", err).Error("Could not get user packages from database")
			return c.SendStatus(fiber.StatusBadRequest)
		}

		packages := make([]Package, 0)
		for _, p := range packs {
			var newPack = Package{
				ID:         uint64(p.ID),
				Tier:       string(p.Tier),
				ValidFrom:  p.ValidFrom,
				ValidUntil: p.ValidUntil,
			}
			packages = append(packages, newPack)
		}

		newUser := User{
			ID:       uint64(user.ID),
			Email:    user.Email,
			Verified: user.VerifiedAt.Valid,
			Packages: packages,
		}
		users = append(users, newUser)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

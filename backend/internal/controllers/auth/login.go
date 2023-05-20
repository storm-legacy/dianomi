package controllers

import (
	"context"
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/jwt"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func Login(c *fiber.Ctx) error {
	var err error

	// * PARSE DATA
	var userData *mod.FormLoginUser
	if err := c.BodyParser(&userData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// * VALIDATE DATA
	userData.Email = strings.ToLower(userData.Email)
	// validate data
	err = mod.Validate.Struct(userData)
	if err != nil {
		log.WithField("err", err).Debug("Could not validate user data")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
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
	queries := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	user, err := queries.GetUserByEmail(ctx, userData.Email)
	if err == sql.ErrNoRows {
		log.WithField("user", userData.Email).Debug("User doesn't exist")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
		})
	}
	if err != nil {
		log.WithField("err", err).Error("SQL query resulted in error")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// * COMPARE PASSWORD
	result, err := argon2.ComparePasswordAndHash(&userData.Password, &user.Password)
	if !result {
		if err != nil {
			log.WithField("err", err).Error("Problem decoding password")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect login information",
		})
	}

	// * GET PACKAGES
	var role string
	userPackages, err := queries.GetPackagesByUserID(ctx, user.ID)
	if err != nil {
		log.WithFields(log.Fields{"user": user, "err": err}).Error("Problem while resolving user packages")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if len(userPackages) < 1 {
		role = "free"
	} else {
		role = string(userPackages[0].Tier)
	}

	// * GENERATE TOKEN WITH CUSTOM INFORMATIONS
	claims := make(map[string]interface{})
	claims["role"] = role
	claims["verified"] = user.VerifiedAt.Valid

	token, err := jwt.GenerateToken(uint64(user.ID), claims)
	if err != nil {
		log.WithField("err", err).Error("Problem generating JWT token")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.WithField("email", userData.Email).Debug("New token generated")
	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data: fiber.Map{
			"token":    token,
			"email":    userData.Email,
			"role":     role,
			"verified": user.VerifiedAt.Valid,
		},
	})
}

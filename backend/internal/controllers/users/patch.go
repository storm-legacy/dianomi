package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	authCtrl "github.com/storm-legacy/dianomi/internal/controllers/auth"
	"github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type PatchUserData struct {
	Email         string    `json:"email" validate:"required,email"`
	Verified      bool      `json:"verified" validate:"boolean"`
	ResetPassword bool      `json:"reset_password" validate:"boolean"`
	Pack          []Package `json:"packages"`
}

func PatchUser(c *fiber.Ctx) error {
	// Get video ID
	idString := c.Params("id")
	paramId, err := strconv.ParseInt(string(idString), 10, 64)
	if err != nil {
		log.WithField("err", err).Debug("ID could not be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Parse sent data
	var data PatchUserData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	// Validate data
	err = models.Validate.Struct(data)
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
	user, err := qtx.GetUserByID(ctx, paramId)
	if err != sql.ErrNoRows && err != nil {
		log.WithField("err", err.Error()).Error("Could not get user from database")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err == sql.ErrNoRows {
		log.WithField("userID", user.ID).Debug("User with specified id doesn't exist")
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Email update
	if err := qtx.UpdateUserEmail(ctx, sqlc.UpdateUserEmailParams{
		ID:    user.ID,
		Email: data.Email,
	}); err != nil {
		log.WithField("userID", user.ID).Error("Users email could not be updated")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Verification update
	if data.Verified && !user.VerifiedAt.Valid {
		now := sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}
		if err := qtx.UpdateUserVerification(ctx, sqlc.UpdateUserVerificationParams{
			ID:         user.ID,
			VerifiedAt: now,
		}); err != nil {
			log.WithField("userID", user.ID).Error("User could not be verified")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	} else if !data.Verified && user.VerifiedAt.Valid {
		nullvalue := sql.NullTime{
			Time:  time.Now(),
			Valid: false,
		}
		if err := qtx.UpdateUserVerification(ctx, sqlc.UpdateUserVerificationParams{
			ID:         user.ID,
			VerifiedAt: nullvalue,
		}); err != nil {
			log.WithField("userID", user.ID).Error("User could not be verified")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	// Trigger password reset
	if data.ResetPassword {
		go authCtrl.AsyncReset(authCtrl.ResetData{
			Email: user.Email,
		})
	}

	// Edit packages
	for _, p := range data.Pack {
		// Verify package
		if p.ValidFrom.Unix() > p.ValidUntil.Unix() {
			log.WithFields(log.Fields{
				"validFrom":  p.ValidFrom,
				"validUntil": p.ValidUntil,
			}).Warn("Time in one or more packages is incorrect")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ValidFrom and ValidUntil values aren't correct",
			})
		}

		// Check if package exists
		packDB, err := qtx.GetPackageByID(ctx, int64(p.ID))
		if err != sql.ErrNoRows && err != nil {
			log.WithField("err", err.Error()).Error("Could not get package from database")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Add new package
		if err == sql.ErrNoRows {
			// Check if package overlaps with another
			items, err := qtx.GetOverlapingPackages(ctx, sqlc.GetOverlapingPackagesParams{
				ValidFrom:  p.ValidFrom,
				ValidUntil: p.ValidUntil,
			})
			if err != nil {
				log.WithField("err", err.Error()).Error("Could not get packages from database")
				return c.SendStatus(fiber.StatusInternalServerError)
			}

			if len(items) > 0 {
				log.New().WithFields(log.Fields{
					"dbPackages": items,
					"newPackage": p,
				}).Warn("New package collides with older one")
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Package collides with another one already present",
				})
			}

			// Add user package
			if err := qtx.GiveUserPackage(ctx, sqlc.GiveUserPackageParams{
				UserID: sql.NullInt64{
					Int64: user.ID,
					Valid: true,
				},
				Tier:       sqlc.Role(p.Tier),
				ValidFrom:  p.ValidFrom,
				ValidUntil: p.ValidUntil,
			}); err != nil {
				log.WithField("err", err.Error()).Error("Could not give user package")
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		} else {
			// Check if package is to be deleted
			if p.Delete {
				if err := qtx.RemoveUserPackage(ctx, packDB.ID); err != nil {
					log.WithField("err", err.Error()).Error("Package could not be removed")
					return c.SendStatus(fiber.StatusInternalServerError)
				}
			}
		}

		// If delete is checked

		// Update package, because it exist
		if err := qtx.UpdatePackage(ctx, sqlc.UpdatePackageParams{
			Tier:       sqlc.Role(p.Tier),
			ValidFrom:  p.ValidFrom,
			ValidUntil: p.ValidUntil,
		}); err != nil {
			log.WithField("err", err.Error()).Error("Package could not be updated")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		log.WithFields(log.Fields{
			"user":    user.ID,
			"package": p,
		}).Info("Package was successfuly updated")

	}

	return c.SendStatus(fiber.StatusOK)
}

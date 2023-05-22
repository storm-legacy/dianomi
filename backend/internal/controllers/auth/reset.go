package controllers

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type PasswordResetData struct {
	Validate       string `json:"validate" validate:"required"`
	Password       string `json:"password" validate:"required,password"`
	PasswordRepeat string `json:"password_repeat" validate:"required,password"`
}

// Function checking if provided UUID is releated to it's task (resetpasword)
func checkResetUUID(code uuid.UUID) error {
	// * CHECK AGANIST DATABASE
	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return err
	}
	queries := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	// Check if is valid
	uuidDb, err := queries.GetVerificationCode(ctx, code)
	if err != nil {
		return err
	}

	// check if isn't used
	if uuidDb.Used {
		return fmt.Errorf("reset code already used. Database value: %t", uuidDb.Used)
	}

	// Check if isn't expired
	now := time.Now().Unix()
	if uuidDb.ValidUntil.Time.Unix() < now {
		return fmt.Errorf("reset code is expired: %d", uuidDb.ValidUntil.Time.Unix())
	}

	// Check if it's reset token
	if uuidDb.TaskType != sqlc.VerifyEmailTypePasswordReset {
		return fmt.Errorf("wrong type of token task: %s", uuidDb.TaskType)
	}

	return nil
}

// Task sending emails
func AsyncReset(resetData ResetData) {
	// * VALIDATE DATA
	resetData.Email = strings.ToLower(resetData.Email)
	// validate data
	err := mod.Validate.Struct(resetData)
	if err != nil {
		log.WithField("err", err).Debug("Could not validate user data")
		return
	}

	// * CHECK AGANIST DATABASE
	// * START(DB BLOCK)
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		log.WithField("err", err).Error("Could not create database connection")
		return
	}
	queries := sqlc.New(db)
	defer db.Close()
	// * END(DB BLOCK)

	user, err := queries.GetUserByEmail(ctx, resetData.Email)
	if err == sql.ErrNoRows {
		log.WithField("user", resetData.Email).Debug("User doesn't exist")
		return
	} else if err != nil {
		log.WithField("err", err).Error("SQL query resulted in error")
		return
	}

	// Create reset code

	code, err := queries.CreateResetCode(ctx, sql.NullInt64{
		Int64: user.ID,
		Valid: true,
	})
	if err != nil {
		log.WithField("error", err.Error()).Error("Could not create reset code")
	}
	frontUrl := config.GetString("APP_FRONT_URL")
	resetUrl := fmt.Sprintf("%s/reset-password?validate=%s", frontUrl, code.Code)

	storagePath := config.GetString("APP_STORAGE_PATH", "./storage")

	// Parse template
	var buf bytes.Buffer
	templatePath := fmt.Sprintf("%s/%s", storagePath, "templates/reset-email.html")
	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		log.WithField("error", err.Error()).Error("Template could not be parsed")
		return
	}
	if err := tmplt.Execute(&buf, struct{ Url string }{Url: resetUrl}); err != nil {
		log.WithField("error", err.Error()).Error("Template could not be recreated")
		return
	}

	if err := sendEmail(Mail{
		To:      []string{resetData.Email},
		Subject: "DianomiTV - Password reset",
		Body:    buf.Bytes(),
	}); err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"url": resetUrl,
		}).Error("Reset email could not be sent.")
		return
	}

	log.Info("Email with reset link send succesfully!")
}

// Route used to generate password reset request
func GenerateReset(c *fiber.Ctx) error {
	// * PARSE DATA
	var resetData *ResetData
	if err := c.BodyParser(&resetData); err != nil {
		log.WithField("error", err.Error()).Debug("Sent data is not correct.")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	go AsyncReset(*resetData)

	log.WithField("email", resetData.Email).Info("User requested password reset!")
	return c.SendStatus(fiber.StatusOK)
}

// Route used to check if token is valid
func GetReset(c *fiber.Ctx) error {
	resetUuidString := c.Query("validate")

	// Check if uuid is correct
	uuidValue, err := uuid.Parse(resetUuidString)
	if err != nil {
		log.WithField("err", err).Debug("Sent UUID is not correct")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := checkResetUUID(uuidValue); err != nil {
		return c.SendStatus(fiber.StatusForbidden)
	}

	return c.SendStatus(fiber.StatusOK)
}

func PostReset(c *fiber.Ctx) error {
	// * PARSE DATA
	var resetData *PasswordResetData
	if err := c.BodyParser(&resetData); err != nil {
		log.WithField("err", err.Error()).Debug("Data send by user, couldn't be parsed")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Check if uuid is correct
	uuidValue, err := uuid.Parse(resetData.Validate)
	if err != nil {
		log.WithField("err", err).Debug("Sent UUID is not correct")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := checkResetUUID(uuidValue); err != nil {
		log.WithField("err", err.Error()).Debug("Sent code couldn't verify the task")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Validate password
	// validate data
	err = mod.Validate.Struct(resetData)
	if err != nil {
		log.WithField("err", err).Debug("Could not validate reset data")
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

	// Invalidate code for later requests
	uuidDb, _ := queries.GetVerificationCode(ctx, uuidValue)
	if err := queries.SetCodeAsUsed(ctx, uuidDb.ID); err != nil {
		log.WithField("err", err).Error("Code could not be set as used")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Encode password
	hashedPassword, err := encodePassword(resetData.Password)
	if err != nil {
		log.WithField("err", err).Error("Password could not be hashed")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Change user password
	if err := queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:       uuidDb.UserID.Int64,
		Password: hashedPassword,
	}); err != nil {
		log.WithField("err", err).Error("User password could not be updated")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	log.WithField("userID", uuidDb.UserID).Info("User password was reset!")
	return c.SendStatus(fiber.StatusOK)
}

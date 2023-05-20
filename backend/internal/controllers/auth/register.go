package controllers

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func generateVerificationCode(userId int64, userEmail string) {
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

	storagePath := config.GetString("APP_STORAGE_PATH", "./storage")
	appUrl := config.GetString("APP_URL")

	// Create verification code and url
	code, err := queries.CreateVerificationCode(ctx, userId)
	if err != nil {
		log.WithField("error", err.Error()).Error("Could not create verification code")
	}
	verificationUrl := fmt.Sprintf("%s/auth/verify?validate=%s", appUrl, code.Code)

	// Skip template generation if smtp is disabled
	smtpEnabled := config.GetBool("APP_SMTP_ENABLED", false)
	if !smtpEnabled {
		log.WithFields(log.Fields{
			"user": userEmail,
			"url":  verificationUrl,
		}).Warn("SMTP is disabled, verfication link here")
		return
	}

	// Parse template
	var buf bytes.Buffer
	templatePath := fmt.Sprintf("%s/%s", storagePath, "templates/verification-email.html")
	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		log.WithField("error", err.Error()).Error("Template could not be parsed")
		return
	}
	if err := tmplt.Execute(&buf, struct{ Url string }{Url: verificationUrl}); err != nil {
		log.WithField("error", err.Error()).Error("Template could not be recreated")
		return
	}

	// Send email with verification code
	smtpHost := config.GetString("APP_SMTP_HOST", "")
	smtpPort := config.GetInt("APP_SMTP_PORT", -1)
	smtpUser := config.GetString("APP_SMTP_USER", "")
	smtpPassword := config.GetString("APP_SMTP_PASSWORD", "")
	smtpNoreplay := config.GetString("APP_SMTP_NOREPLAY", "")

	// smtpTLS := config.GetString("APP_SMTP_TLS", "")
	smtpHostAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	toEmails := []string{userEmail}
	msg := buildMessage(mod.Mail{
		Sender:  smtpNoreplay,
		To:      toEmails,
		Subject: "Verify your DianomiTV email address",
		Body:    buf.String(),
	})

	if err := smtp.SendMail(smtpHostAddr, auth, smtpNoreplay, toEmails, msg); err != nil {
		log.WithField("error", err.Error()).Error("Email could not be sent")
		return
	}

	log.Info("Email with verification code sent successfuly!")
}

func Register(c *fiber.Ctx) error {
	var err error

	// * PARSE DATA
	var userData *mod.FormRegisterUser
	if err := c.BodyParser(&userData); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// * VALIDATE DATA
	userData.Email = strings.ToLower(userData.Email)
	userData.Email = strings.TrimSpace(userData.Email)
	// validate data
	err = mod.Validate.Struct(userData)
	if err != nil {
		log.WithField("err", err).Debug("Could not validate user data")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Incorrect registration information provided",
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

	// Check if exists
	_, err = queries.GetUserByEmail(ctx, userData.Email)
	if err != sql.ErrNoRows {
		if err != nil {
			log.WithField("err", err).Error("SQL query resulted in error")
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		log.WithField("user", userData.Email).Debug("User already exists")
		return c.Status(fiber.StatusBadRequest).JSON(mod.Response{
			Status: "error",
			Data:   "Account with this e-mail address already exists",
		})
	}

	// Encode password
	hashedPassword, err := argon2.EncodePassword(&userData.Password)
	if err != nil {
		log.WithField("err", err).Error("Password could not be hashed")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Insert user to database
	user, err := queries.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Email:    userData.Email,
			Password: hashedPassword,
		})
	if err != nil {
		log.WithField("err", err).Error("Account could not be created")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	go generateVerificationCode(user.ID, user.Email)

	log.WithField("email", userData.Email).Debug("New account was created")
	return c.Status(fiber.StatusOK).JSON(mod.Response{
		Status: "success",
		Data:   "Account was created",
	})
}

package controllers

import (
	"errors"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
)

type ResetData struct {
	Email string `json:"email" validate:"required,email"`
}

type TemplateData struct {
	Url  string
	Code uuid.UUID
}

// https://zetcode.com/golang/email-smtp/
type Mail struct {
	// Sender  string
	To      []string
	Subject string
	Body    []byte
}

func SendEmail(mail Mail) error {
	// Skip template generation if smtp is disabled
	smtpEnabled := config.GetBool("APP_SMTP_ENABLED", false)
	if !smtpEnabled {
		return errors.New("smtp is disabled")
	}

	// Send email with verification code
	smtpHost := config.GetString("APP_SMTP_HOST", "")
	smtpPort := config.GetInt("APP_SMTP_PORT", -1)
	smtpUser := config.GetString("APP_SMTP_USER", "")
	smtpPassword := config.GetString("APP_SMTP_PASSWORD", "")
	smtpNoreplay := config.GetString("APP_SMTP_NOREPLAY", "")
	smtpHostAddr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	// smtpTLS := config.GetString("APP_SMTP_TLS", "")

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", smtpNoreplay)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	if err := smtp.SendMail(smtpHostAddr, auth, smtpNoreplay, mail.To, []byte(msg)); err != nil {
		log.WithField("error", err.Error()).Error("")
		return errors.New("email could not be sent")
	}

	return nil
}

func encodePassword(password string) (string, error) {
	result, err := argon2.EncodePassword(&password)
	if err != nil {
		log.WithField("err", err).Error("Password could not be hashed")
		return "", err
	}
	return result, nil
}

func comparePasswords(password string, hashedPassword string) (bool, error) {
	result, err := argon2.ComparePasswordAndHash(&password, &hashedPassword)
	if err != nil {
		return false, err
	}
	return result, nil
}

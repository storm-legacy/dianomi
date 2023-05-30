package controllers

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	controllers "github.com/storm-legacy/dianomi/internal/controllers/auth"
	mod "github.com/storm-legacy/dianomi/internal/models"
	"github.com/storm-legacy/dianomi/pkg/config"
)

type Error_Reports struct {
	ErrorTitle       string `json:"ErrorTitle" validate:"required"`
	ErrorDescription string `json:"ErrorDescription" validate:"required"`
	ReportedBy       string `json:"ReportedBy" validate:"required"`
}

func PostReport(c *fiber.Ctx) error {
	var data Error_Reports
	if err := c.BodyParser(&data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := mod.Validate.Struct(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("Data validation error (%s)", err.Error()),
		})
	}

	storagePath := config.GetString("APP_STORAGE_PATH", "./storage")
	contactEmail := config.GetString("APP_REPORT_EMAIL", "")

	// Parse template
	var buf bytes.Buffer
	templatePath := fmt.Sprintf("%s/%s", storagePath, "templates/report-email.html")
	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		log.WithField("error", err.Error()).Error("Template could not be parsed")
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if err := tmplt.Execute(&buf, struct {
		ErrorTitl        string
		ErrorDescription string
		ReportedBy       string
	}{ErrorTitl: data.ErrorTitle,
		ErrorDescription: data.ErrorDescription,
		ReportedBy:       data.ReportedBy}); err != nil {
		log.WithField("error", err.Error()).Error("Template could not be recreated")
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	message := controllers.Mail{
		To:      []string{contactEmail},
		Subject: "Problem report",
		Body:    buf.Bytes(),
	}

	if err := controllers.SendEmail(message); err != nil {
		log.WithFields(log.Fields{
			"err":     err.Error(),
			"subject": "Problem report",
			"content": data.ErrorDescription,
			"author":  data.ReportedBy,
		}).Error("Reset email could not be sent.")
		c.SendStatus(fiber.StatusOK)
	}

	return c.SendStatus(fiber.StatusAccepted)
}

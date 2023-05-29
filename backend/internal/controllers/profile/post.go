package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
)

func PostPayment(c *fiber.Ctx) error {

	id := c.Locals("sub").(uint64)

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

	user, err := qtx.GetUserByID(ctx, int64(id))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	stripe.Key = config.GetString("STRIPE_API_KEY")
	frontURL := config.GetString("APP_FRONT_URL")
	productID := config.GetString("STRIPE_PRODUCT_ID")

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(productID),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:          stripe.String("payment"),
		SuccessURL:    stripe.String(frontURL + "/payment?status=success"),
		CancelURL:     stripe.String(frontURL + "/payment?status=canceled"),
		CustomerEmail: &user.Email,
	}

	s, _ := session.New(params)
	fmt.Println(s.Status)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": s.URL,
	})
}

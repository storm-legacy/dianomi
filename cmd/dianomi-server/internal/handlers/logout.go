package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/jwt"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

func LogoutUser(token string) (result Result) {
	// Database block (start)
	// Opening database connection
	db, err := sql.Open("postgres", config.GetString("PG_CONNECTION_STRING"))
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}
	ctx := context.Background()
	queries := sqlc.New(db)
	defer db.Close()
	// Database block (end)

	// Parse token into claims
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}
	userid := int64(claims["sub"].(float64))
	valid := time.Unix(int64(claims["exp"].(float64)), 0)

	// Check if token was already revoked
	_, err = queries.CheckToken(ctx, token)
	if err != sql.ErrNoRows {
		if err != nil {
			return Result{fiber.StatusBadRequest, err.Error(), nil}
		}
		return Result{fiber.StatusBadRequest, "token already revoked", nil}
	}

	// Check if token already expired
	if time.Now().Unix() > valid.Unix() {
		return Result{fiber.StatusBadRequest, "token already expired", nil}
	}

	// Add token to revoked list
	err = queries.AddRevokedToken(ctx, sqlc.AddRevokedTokenParams{
		Token:      token,
		UserID:     userid,
		ValidUntil: time.Time(valid),
	})
	if err != nil {
		fmt.Print(err)
	}

	// Success
	return Result{fiber.StatusOK, "", fiber.Map{"status": "logged out"}}
}

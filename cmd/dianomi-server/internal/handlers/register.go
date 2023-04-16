package handlers

import (
	"context"
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type FormRegisterUser struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	PasswordRepeated string `json:"password-repeat"`
}

func RegisterUser(userData *FormRegisterUser) (result Result) {
	// Change email to lowercase
	userData.Email = strings.ToLower(userData.Email)

	// Check if email and password are valid
	if !emailRegex.MatchString(userData.Email) {
		return Result{fiber.StatusBadRequest, "email is invalid", nil}
	}

	if len(userData.Password) < 8 {
		return Result{fiber.StatusBadRequest, "password is too simple", nil}
	}

	// Check if password is the same as repeated one
	if userData.Password != userData.PasswordRepeated {
		return Result{fiber.StatusBadRequest, "passwords are not the same", nil}
	}

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

	// Get the user from database
	_, err = queries.GetUserByEmail(ctx, userData.Email)
	if err == nil {
		return Result{fiber.StatusBadRequest, "account with this username already exists", nil}

	} else if err != sql.ErrNoRows {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}

	// Encode password
	hashedPassword, err := argon2.EncodePassword(&userData.Password)
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}

	// Insert user to database
	err = queries.CreateUser(
		ctx,
		sqlc.CreateUserParams{
			Email:    userData.Email,
			Password: hashedPassword,
		})
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}

	// Success
	return Result{fiber.StatusOK, "", fiber.Map{"status": "success"}}
}

package handlers

import (
	"context"
	"database/sql"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/storm-legacy/dianomi/pkg/argon2"
	"github.com/storm-legacy/dianomi/pkg/config"
	"github.com/storm-legacy/dianomi/pkg/jwt"
	"github.com/storm-legacy/dianomi/pkg/sqlc"
)

type FormLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(loginUser *FormLoginUser) (result Result) {
	// Change email to lowercase
	loginUser.Email = strings.ToLower(loginUser.Email)

	// Check if email and password are valid
	if !emailRegex.MatchString(loginUser.Email) {
		return Result{fiber.StatusBadRequest, "email is invalid", nil}
	}

	if len(loginUser.Password) < 8 {
		return Result{fiber.StatusBadRequest, "password is invalid", nil}
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
	user, err := queries.GetUserByEmail(ctx, loginUser.Email)
	if err == sql.ErrNoRows {
		return Result{fiber.StatusBadRequest, "user or password is incorrect", nil}

	} else if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}

	// Compare password of the user
	isTheSame, err := argon2.ComparePasswordAndHash(&loginUser.Password, &user.Password)
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}

		// Password is incorrect
	} else if !isTheSame {
		return Result{fiber.StatusBadRequest, "user or password is incorrect", nil}
	}

	// Get user packages
	var role string
	pack, err := queries.GetPackagesByUserID(ctx, user.ID)
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}
	if len(pack) < 1 {
		role = "free"
	} else {
		role = string(pack[0].Tier)
	}

	token, err := jwt.GenerateToken(user.ID, user.Email, role)
	if err != nil {
		return Result{fiber.StatusInternalServerError, err.Error(), nil}
	}

	return Result{fiber.StatusOK, "", fiber.Map{"token": token}}
}

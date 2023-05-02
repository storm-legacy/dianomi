package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate = validator.New()
)

func passwordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// At least 8 characters
	if len(password) < 8 {
		return false
	}

	// At least 1 uppercase letter
	if ok, _ := regexp.MatchString(`[A-Z]`, password); !ok {
		return false
	}

	// At least 1 lowercase letter
	if ok, _ := regexp.MatchString(`[a-z]`, password); !ok {
		return false
	}

	// At least 1 digit
	if ok, _ := regexp.MatchString(`[0-9]`, password); !ok {
		return false
	}

	// At least 1 special character
	if ok, _ := regexp.MatchString(`[#?!@$%^&*-]`, password); !ok {
		return false
	}

	return true
}

func init() {
	Validate.RegisterValidation("password", passwordValidation)
}

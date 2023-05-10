package models

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate = validator.New()
	tagRegex *regexp.Regexp      = regexp.MustCompilePOSIX(`^[a-z]{3,12}$`)
)

func tagsValidation(fl validator.FieldLevel) bool {
	tags := fl.Field().Interface().([]string)
	// Limit tags number to 10
	if len(tags) > 10 {
		return false
	}

	// Trim tags
	for _, tag := range tags {
		if !tagRegex.MatchString(tag) {
			return false
		}
	}

	return true
}

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
	Validate.RegisterValidation("tags", tagsValidation)
}

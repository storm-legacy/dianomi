package handlers

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

const (
	SUCCESS      int = 2
	CLIENT_ERROR int = 4
	SERVER_ERROR int = 5
	OTHER_ERROR  int = 9
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Result struct {
	StatusCode   int
	ErrorMessage string
	Data         fiber.Map
}

func (r *Result) Status() int {
	firstDigit := r.StatusCode / 100

	switch firstDigit {
	case SUCCESS:
		return SUCCESS

	case CLIENT_ERROR:
		return CLIENT_ERROR

	case SERVER_ERROR:
		return SERVER_ERROR

	default:
		return OTHER_ERROR
	}
}

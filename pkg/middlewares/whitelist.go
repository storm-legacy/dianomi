package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

var (
	whitelist = []string{"127.0.0.1"}
)

func contains(array *[]string, value string) bool {
	for _, v := range *array {
		if v == value {
			return true
		}
	}
	return false
}

func WhitelistMiddleware(c *fiber.Ctx) error {
	if contains(&whitelist, c.IP()) {
		return c.Next()
	}
	return c.SendStatus(fiber.StatusForbidden)
}

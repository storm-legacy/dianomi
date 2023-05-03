package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func DebugMiddleware(c *fiber.Ctx) error {
	// Request
	// Print the request headers
	fmt.Println("Headers:")
	fmt.Println(string(c.Request().Header.Header()))

	// Print the request body
	fmt.Println("Body:")
	fmt.Println(string(c.Request().Body()))

	return c.Next()
}

package routes

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
)

// Utility functions
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Information about service status
func Healthcheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"healthy": true})
}

// Some informations about service usage
func HealthStats(c *fiber.Ctx) error {
	if c.IP() == "127.0.0.1" {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"healthy": true,
			"resources": fiber.Map{
				"allocMb":      bToMb(m.Alloc),
				"totalAllocMb": bToMb(m.TotalAlloc),
				"sysMb":        bToMb(m.Sys),
				"numGCMb":      bToMb(uint64(m.NumGC)),
			},
		})
	}
	return c.SendStatus(fiber.StatusForbidden)
}

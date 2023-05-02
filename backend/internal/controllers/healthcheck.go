package controllers

import (
	"runtime"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

var (
	whitelist = []string{"127.0.0.1"}
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
func InternalHealthcheck(c *fiber.Ctx) error {
	// Whitelist only for localhost
	if !slices.Contains(whitelist, c.IP()) {
		return c.SendStatus(fiber.StatusNotFound)
	}
	// Calculate application data usage
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

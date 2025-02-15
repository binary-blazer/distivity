package routes

import (
	"github.com/gofiber/fiber/v2"
)

func StatusHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "API is running",
		"success": true,
	})
}

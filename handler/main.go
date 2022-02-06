package handler

import "github.com/gofiber/fiber/v2"

func DocImport(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  1,
		"message": "...",
		"data":    nil,
	})
}

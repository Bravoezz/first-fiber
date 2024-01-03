package middleware

import "github.com/gofiber/fiber/v2"

func ErrorMiddleware(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return nil
}
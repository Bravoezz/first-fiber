package server

import (
	"fmt"

	"github.com/Bravoezz/first-fiber/middleware"
	"github.com/Bravoezz/first-fiber/modules"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Start(address string) {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: middleware.ErrorMiddleware,
	})

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server online ✅")
	})

	// setup router
	modules.MainRouter(app.Group("/api/v1"))

	app.Use(recover.New())

	app.Listen(fmt.Sprintf(":%s",address))
}
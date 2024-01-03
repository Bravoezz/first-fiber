package modules

import (
	"github.com/Bravoezz/first-fiber/modules/tasks"
	"github.com/gofiber/fiber/v2"
)

func MainRouter(apiV1 fiber.Router) {

	// task routes
	tasks.TaskRouter(apiV1.Group("/task"))

}
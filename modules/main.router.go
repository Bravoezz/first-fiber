package modules

import (
	"github.com/Bravoezz/first-fiber/modules/tasks"
	"github.com/Bravoezz/first-fiber/modules/users"
	"github.com/gofiber/fiber/v2"
)

func MainRouter(apiV1 fiber.Router) {

	tasks.TaskRouter(apiV1.Group("/task"))
	users.UserRouter(apiV1.Group("/user"))
}
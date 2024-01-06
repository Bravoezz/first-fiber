package tasks

import (
	"github.com/Bravoezz/first-fiber/modules/models"
	"github.com/gofiber/fiber/v2"
)

type ITaskController interface {
	GetAllTasks() (func(c *fiber.Ctx) error)
}

type ITaskService interface {
	GetAllTasks() ([]models.Task, error)
}

type ITaskRepository interface {
	GetAll() ([]models.Task, error)
}

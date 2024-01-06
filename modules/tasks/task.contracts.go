package tasks

import (
	"github.com/Bravoezz/first-fiber/modules/models"
	"github.com/gofiber/fiber/v2"
)

type ITaskController interface {
	GetAllTasks(c *fiber.Ctx) error
	GetTaskById(c *fiber.Ctx) error
}

type ITaskService interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int) (models.Task, error)
}

type ITaskRepository interface {
	GetAll() ([]models.Task, error)
	GetById(id int) (models.Task, error)
}

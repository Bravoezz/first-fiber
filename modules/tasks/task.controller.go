package tasks

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct{
	taskService ITaskService
}

func (tsk TaskController) GetAllTasks() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tasks, err := tsk.taskService.GetAllTasks()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return err
		}

		return c.Status(200).JSON(tasks)
	}
}

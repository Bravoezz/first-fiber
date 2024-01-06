package tasks

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskController struct {
	taskService ITaskService
}

func NewController(s ITaskService) ITaskController {
	return &TaskController{
		taskService: s,
	}
}

func (tsk TaskController) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := tsk.taskService.GetAllTasks()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	return c.Status(200).JSON(tasks)

}

func (tsk TaskController) GetTaskById(c *fiber.Ctx) error {
	id,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	task, err := tsk.taskService.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("error task ctrll")
	}

	return c.Status(200).JSON(task)
}

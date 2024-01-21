package tasks

import (
	"fmt"
	"runtime"
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

// GetEspecifyTask implements ITaskController.
func (tsk TaskController) GetEspecifyTask(c *fiber.Ctx) error {
	
	taskData, err := tsk.taskService.GetEspecifyTask()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"res": true,
		"data": *taskData,
	})
}

func (tsk TaskController) GetAllTasks(c *fiber.Ctx) error {
	fmt.Println("Numeros CPUs: ", runtime.NumCPU())

	tasks, err := tsk.taskService.GetAllTasks()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	// comprobando que la direccion de memoria de slice task en controller sea el mismo que en repository
	fmt.Printf("Controller: %p\n", tasks)

	return c.Status(200).JSON(tasks)
}

func (tsk TaskController) GetTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	task, err := tsk.taskService.GetTaskById(id)
	if err != nil {
		return fmt.Errorf("error task ctrll")
	}

	return c.Status(200).JSON(task)
}

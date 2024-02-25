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

func TaskAllResponse(c *fiber.Ctx,status int,res bool ,data interface{}) error {
	return c.Status(status).JSON(fiber.Map{ "res": res, "data": data}) 
}

// GetEspecifyTask implements ITaskController.
func (tsk TaskController) GetEspecifyTask(c *fiber.Ctx) error {
	
	taskData, err := tsk.taskService.GetEspecifyTask()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}

	return TaskAllResponse(c,200,true,*taskData)
}

func (tsk TaskController) GetAllTasks(c *fiber.Ctx) error {
	fmt.Println("Numeros CPUs: ", runtime.NumCPU())

	tasks, err := tsk.taskService.GetAllTasks()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return TaskAllResponse(c,404,false,struct{}{})
	}

	// comprobando que la direccion de memoria de slice task en controller sea el mismo que en repository
	fmt.Printf("Controller: %p\n", tasks)

	return TaskAllResponse(c,200,true,tasks)
}

func (tsk TaskController) GetTaskById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	task, err := tsk.taskService.GetTaskById(id)
	if err != nil {
		return TaskAllResponse(c,404,false,struct{}{})
	}

	return TaskAllResponse(c,200,true,task)
}

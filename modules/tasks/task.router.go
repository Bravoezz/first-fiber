package tasks

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func routine(done chan TaskResponse) {
	fmt.Println("prepaando tarea")
	task := TaskResponse{1, "Hacer la comida"}

	time.Sleep(time.Millisecond * time.Duration(500))
	fmt.Println("tarea terminada")
	done <- task
}

type TaskResponse struct {
	Id   int
	Name string
}

//** tre tipos de chanels o canales
//** bidireccionales bidireccionales
//** receive-only	solo reciven
//** send-only		solo envian


func TaskRouter(router fiber.Router) {
	// var taskController ITaskController = TaskController{taskService: TaskService{taskRepository: TaskRepository{}}}
	taskRepository := Newrepository()
	taskService := NewService(taskRepository)
	taskController := NewController(taskService)

	router.Get("/all", taskController.GetAllTasks)
	router.Get("/:id", taskController.GetTaskById)
	
	// aqui se usar un chanel de tipo bidireccional
	router.Get("/go-routine", func (c *fiber.Ctx) error {

		done := make(chan TaskResponse)

		go routine(done)

		resp := <- done
		close(done)

		return c.Status(200).JSON(resp)
	})
}

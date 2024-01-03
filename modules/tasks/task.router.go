package tasks

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
)

func routine(done chan TaskResponse) {
	fmt.Println("prepaando tarea")
	task := TaskResponse{ 1,"Hacer la comida"}

	time.Sleep(time.Millisecond * time.Duration(500))
	fmt.Println("tarea terminada")
	done <- task
}

type TaskResponse struct {
	Id   int
	Name string
}

func TaskRouter(router fiber.Router) {
	router.Get("/all", func(c *fiber.Ctx) error {

		done := make(chan TaskResponse)

		go routine(done)
		resp := <- done
		close(done)

		return c.Status(200).JSON(resp)
	})
}

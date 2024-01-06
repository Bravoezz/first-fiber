package tasks

import "github.com/Bravoezz/first-fiber/modules/models"

type TaskService struct {
	taskRepository ITaskRepository
}

func (tsr TaskService) GetAllTasks() ([]models.Task, error) {

	return tsr.taskRepository.GetAll()
}

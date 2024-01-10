package tasks

import "github.com/Bravoezz/first-fiber/modules/models"


type TaskService struct {
	taskRepository ITaskRepository
}

func NewService(r ITaskRepository) ITaskService {
	return &TaskService{
		taskRepository: r,
	}
}

func (tsr TaskService) GetAllTasks() (*[]models.Task, error) {
	return tsr.taskRepository.GetAll()
}

func (tsr TaskService) GetTaskById(id int) (models.Task, error)  {
	return tsr.taskRepository.GetById(id)
}

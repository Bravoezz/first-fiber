package tasks

import (
	"fmt"

	"github.com/Bravoezz/first-fiber/db"
	"github.com/Bravoezz/first-fiber/modules/models"
)

type TaskRepository struct{}

func Newrepository() ITaskRepository { return &TaskRepository{} }

func (trp TaskRepository) GetAll() (*[]models.Task, error) {
	var tasks []models.Task

	err := db.DB.Where(&models.Task{UserId: 1}).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	// comprobando que la direccion de memoria de slice task en repository sea el mismo que en controller
	fmt.Printf("Repository: %p\n", &tasks)

	return &tasks, nil
}

func (trp TaskRepository) GetById(id int) (models.Task, error) {
	var task models.Task

	err := db.DB.First(&task,id).Error
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}


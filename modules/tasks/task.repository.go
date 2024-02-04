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

	// err := db.DB.Where(&models.Task{UserId: 1}).Find(&tasks).Error //! se filtra todos las task de el usuario con id 1
	err := db.DB.Where(&models.Task{}).Find(&tasks).Error
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
		fmt.Println("hola", err.Error()) //!en caso de no encontrar devuelve este mensaje en el error -> record not found
		return models.Task{}, err
	}

	return task, nil
}


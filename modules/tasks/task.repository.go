package tasks

import (
	"github.com/Bravoezz/first-fiber/db"
	"github.com/Bravoezz/first-fiber/modules/models"
)

type TaskRepository struct{}

func (trp TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task

	err := db.DB.Where(&models.Task{UserId: 1}).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

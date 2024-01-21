package tasks

import (
	"fmt"
	"sync"

	"github.com/Bravoezz/first-fiber/modules/models"
)

type TaskService struct {
	taskRepository ITaskRepository
}


type EspecifyTask struct {
	Task3 models.Task
	Task4 models.Task
}

func NewService(r ITaskRepository) ITaskService {
	return &TaskService{
		taskRepository: r,
	}
}

func (tsr TaskService) GetAllTasks() (*[]models.Task, error) {
	return tsr.taskRepository.GetAll()
}

func (tsr TaskService) GetTaskById(id int) (models.Task, error) {
	return tsr.taskRepository.GetById(id)
}

func (tsr TaskService) getTaskConc3(wg *sync.WaitGroup, data *EspecifyTask)  {
	var err error
	data.Task3, err = tsr.taskRepository.GetById(3)
	if err != nil {
		fmt.Println("error en 3", err.Error())
	}
	wg.Done()
}

func (tsr TaskService) getTaskConc4(wg *sync.WaitGroup, data *EspecifyTask)  {
	var err error
	data.Task4, err = tsr.taskRepository.GetById(4)
	if err != nil {
		fmt.Println("error en 4", err.Error())
	}
	wg.Done()
}
// GetEspecifyTask implements ITaskService.
// func (tsr TaskService) GetEspecifyTask() (*EspecifyTask, error)  {
// 	wg := sync.WaitGroup{}
// 	wg.Wait()
// 	taskData := EspecifyTask{}

// 	taskData.Task3, _ = tsr.taskRepository.GetById(3)
// 	taskData.Task4, _ = tsr.taskRepository.GetById(4)

// 	return &taskData, nil
// }

// !without concurrenci
func (tsr TaskService) GetEspecifyTask() (*EspecifyTask, error)  {
	wg := sync.WaitGroup{}
	taskData := EspecifyTask{}

	wg.Add(2)
	go tsr.getTaskConc3(&wg, &taskData)
	go tsr.getTaskConc4(&wg, &taskData)
	
	wg.Wait()
	// fmt.Printf("dat3: %+v\n",taskData.Task3)
	// fmt.Printf("dat4: %+v\n",taskData.Task4)

	return &taskData, nil
}
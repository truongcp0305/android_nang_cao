package service

import (
	"android-service/model"
	"android-service/usecase/repository"
	"encoding/json"
	"errors"
)

type TaskService struct {
	database repository.Database
}

func NewTaskService(db repository.Database) *TaskService {
	return &TaskService{
		database: db,
	}
}

func (s *TaskService) DetailTask(task *model.Task) error {
	err := s.database.GetTaskById(task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) GetList(task *model.Task) ([]model.Task, error) {
	tasks, err := s.database.GetListTaskByUserId(task)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *TaskService) CreateTask(data string) error {
	task := model.Task{}
	err := json.Unmarshal([]byte(data), &task)
	if err != nil {
		return errors.New("Invalid data to create task")
	}
	err = s.database.CreateTask(&task)
	if err != nil {
		return err
	}
	return nil
}

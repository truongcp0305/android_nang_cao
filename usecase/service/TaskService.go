package service

import (
	"android-service/model"
	"android-service/usecase/repository"
	"encoding/json"
	"errors"
	"math/rand"
	"time"
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
	task.Id = generateUUID()
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

func (s *TaskService) UpdateTask(data string) error {
	task := model.Task{}
	err := json.Unmarshal([]byte(data), &task)
	if err != nil {
		return errors.New("Invalid data to create task")
	}
	err = s.database.UpdateTask(&task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(task *model.Task) error {
	err := s.database.DeleteTask(task)
	return err
}

func generateUUID() string {
	// Sử dụng thời gian hiện tại để tạo một seed ngẫu nhiên
	rand.Seed(time.Now().UnixNano())

	// Tạo một chuỗi ngẫu nhiên có độ dài 32 ký tự
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	uuid := make([]byte, 32)
	for i := range uuid {
		uuid[i] = charset[rand.Intn(len(charset))]
	}

	// Thêm một dấu gạch ngang vào vị trí thứ 8 và 13 để tạo định dạng UUID
	uuid[8] = '-'
	uuid[13] = '-'

	return string(uuid)
}

package repository

import "android-service/model"

type Database interface {
	GetTaskById(*model.Task) error
	GetListTaskByUserId(task *model.Task) ([]model.Task, error)
	CreateTask(task *model.Task) error
}

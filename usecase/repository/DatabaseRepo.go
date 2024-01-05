package repository

import "android-service/model"

type Database interface {
	GetTaskById(task *model.Task) (*model.Task, error)
	GetListTaskByUserId(task *model.Task) ([]model.Task, error)
	CreateTask(task *model.Task) error
	DeleteTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	CreateUser(user *model.User) error
	GetListUser() ([]model.User, error)
	GetUserByName(user *model.User) (*model.User, error)
	GetUserByUserNameAndPass(user *model.User) (*model.User, error)
	CreateUserInfo(info *model.UserInfo) error
	GetUserInfo(info *model.UserInfo) (*model.UserInfo, error)
	UpdateUserInfo(info *model.UserInfo) error
	UpdateUser(user *model.User) error
	GetAllTask() ([]model.Task, error)
	GetUserTasks(user model.User) (int, error)
	GetDoneTasks(user model.User) (int, error)
	GetAssignTasks(task model.Task) ([]model.Task, error)
}

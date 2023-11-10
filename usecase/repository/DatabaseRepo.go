package repository

import "android-service/model"

type Database interface {
	GetTaskById(*model.Task) error
	GetListTaskByUserId(task *model.Task) ([]model.Task, error)
	CreateTask(task *model.Task) error
	DeleteTask(task *model.Task) error
	UpdateTask(task *model.Task) error
	CreateUser(user *model.User) error
	GetListUser() ([]model.User, error)
	GetUserByName(user *model.User) error
	GetUserByUserNameAndPass(user *model.User) error
	InsertWord(words []model.Word) error
	GetWordsForQuestion(level string) ([]model.Word, error)
	CreateUserInfo(info *model.UserInfo) error
	GetUserInfo(info *model.UserInfo) error
	UpdateUserInfo(info *model.UserInfo) error
}

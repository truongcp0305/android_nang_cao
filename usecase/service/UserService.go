package service

import (
	"android-service/model"
	"android-service/usecase/repository"
	"errors"
)

type UserService struct {
	database repository.Database
}

func NewUserService(database repository.Database) *UserService {
	return &UserService{
		database: database,
	}
}

func (us *UserService) CreateAccount(user *model.User) (string, error) {
	err := us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		user.UserId = generateUUID()
		err = us.database.CreateUser(user)
		if err != nil {
			return "", err
		}
		return user.UserId, nil
	}
	return "", errors.New("username already exists")
}

func (us *UserService) Login(user *model.User) (string, error) {
	err := us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}

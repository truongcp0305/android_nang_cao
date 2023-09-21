package service

import (
	"android-service/model"
	"android-service/usecase/repository"
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
	user.UserId = generateUUID()
	err := us.database.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}

func (us *UserService) Login(user *model.User) (string, error) {
	err := us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}

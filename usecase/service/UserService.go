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

func (us *UserService) CreateAccount(user *model.User) (model.UserInfo, error) {
	err := us.database.GetUserByName(user)
	if err != nil {
		user.UserId = generateUUID()
		err = us.database.CreateUser(user)
		if err != nil {
			return model.UserInfo{}, err
		}
		info := model.UserInfo{
			UserId:   user.UserId,
			UserName: user.UserName,
			Point:    "0",
		}
		err = us.database.CreateUserInfo(&info)
		if err != nil {
			return model.UserInfo{}, err
		}
		return info, nil
	}
	return model.UserInfo{}, errors.New("username already exists")
}

func (us *UserService) Login(user *model.User) (model.UserInfo, error) {
	err := us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return model.UserInfo{}, err
	}
	info := model.UserInfo{
		UserId: user.UserId,
	}
	err = us.database.GetUserInfo(&info)
	if err != nil {
		return model.UserInfo{}, err
	}
	return info, nil
}

func (us *UserService) Updateinfor(user *model.UserInfo) error {
	err := us.database.UpdateUserInfo(user)
	return err
}

func (us *UserService) GetList() ([]model.User, error) {
	users, err := us.database.GetListUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

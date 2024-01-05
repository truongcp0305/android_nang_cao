package service

import (
	"android-service/crypt"
	"android-service/model"
	"android-service/usecase/repository"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
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
	_, err := us.database.GetUserByName(user)
	if err != nil {
		user.UserId = generateUUID()
		user.Try = 5
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
	uByName := model.User{
		UserName: user.UserName,
	}
	ub, err := us.database.GetUserByName(&uByName)
	if err != nil {
		return model.UserInfo{}, err
	}
	if ub.Try == 0 {
		t, err := time.Parse(time.RFC3339, ub.UnlockTime)
		if err != nil {
			return model.UserInfo{}, err
		}
		if t.Before(time.Now()) {
			ub.Try = 5
			ub.Lock = false
		} else {
			return model.UserInfo{}, fmt.Errorf("You have been lock for 30 minutes")
		}
	}
	if ub.Pass != user.Pass {
		ub.Try--
		ub.UnlockTime = time.Now().Add(30 * time.Minute).Format(time.RFC3339)
		err := us.database.UpdateUser(ub)
		if err != nil {
			return model.UserInfo{}, err
		}
		return model.UserInfo{}, fmt.Errorf("Invalid Password; try: %d", ub.Try)
	}
	user, err = us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return model.UserInfo{}, err
	}
	info := model.UserInfo{
		UserId: user.UserId,
	}
	i, err := us.database.GetUserInfo(&info)
	if err != nil {
		return model.UserInfo{}, err
	}
	user.Try = 5
	go us.database.UpdateUser(user)
	return *i, nil
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
	for i := range users {
		t, err := us.database.GetUserTasks(users[i])
		if err != nil {
			return nil, err
		}
		users[i].TotalTask = t
		done, err := us.database.GetDoneTasks(users[i])
		if err != nil {
			return nil, err
		}
		users[i].TotalTaskDone = done
		ti, _ := time.Parse(time.RFC3339, users[i].UnlockTime)
		if !ti.Before(time.Now()) {
			users[i].Lock = true
		} else {
			users[i].Lock = false
		}
	}
	return users, nil
}

func (us *UserService) GetAssignTasks(userId string) ([]model.Task, error) {
	task := model.Task{
		UserId: userId,
	}
	tasks, err := us.database.GetAssignTasks(task)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

const (
	emailAddr = "boynhabecp@gmail.com"
	emialPass = "okqq ifeb wwza izau"
)

func (us *UserService) SendMailResetPass(mailAddr string) error {
	user := model.User{
		UserName: mailAddr,
	}
	_, err := us.database.GetUserByName(&user)
	if err != nil {
		return errors.New("user do not exist!")
	}
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailAddr)
	mailer.SetHeader("To", mailAddr)
	mailer.SetHeader("Subject", "Reset Password")
	rand.Seed(time.Now().UnixNano())
	newPass := int64(rand.Intn(900000) + 100000)
	data, _ := createResetData(mailAddr, strconv.FormatInt(newPass, 10))
	resetLink := "https://android-service.fly.dev/reset/" + data
	body := fmt.Sprintf("<h1>Click the link to reset your password:</h1> "+
		"<p>Your new password will be: %d<p>"+
		"<p><a href=\"%s\">Reset Password</a></p>", newPass, resetLink)
	mailer.SetBody("text/html", body)
	dialer := gomail.NewDialer("smtp.gmail.com", 587, emailAddr, emialPass)
	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}
	return nil
}

func createResetData(email string, newPass string) (string, error) {
	expired := time.Now().Add(30 * time.Minute).Unix()
	data := email + "|" + newPass + "|" + strconv.FormatInt(expired, 10)
	en, err := crypt.Encrypt(data)
	if err != nil {
		return "", err
	}
	return en, nil
}

func (us *UserService) ResetPass(data string) error {
	v, err := crypt.Decrypt(data)
	if err != nil {
		return err
	}
	val := strings.Split(v, "|")
	if len(val) < 3 {
		return errors.New("invalid data")
	}
	sec, err := strconv.Atoi(val[2])
	if err != nil {
		return err
	}
	t := time.Unix(int64(sec), 0)
	if t.Before(time.Now()) {
		return errors.New("reset link expried!")
	}
	user := model.User{
		UserName: val[0],
	}
	u, err := us.database.GetUserByName(&user)
	if err != nil {
		return err
	}
	u.Pass = val[1]
	u.Try = 5
	err = us.database.UpdateUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UpdatePass(user *model.User, newPass string) error {
	user, err := us.database.GetUserByUserNameAndPass(user)
	if err != nil {
		return err
	}
	user.Pass = newPass
	err = us.database.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) Lock(user *model.User) error {
	u, err := us.database.GetUserByName(user)
	if err != nil {
		return nil
	}
	u.UnlockTime = time.Now().Add(1000000 * time.Hour).Format(time.RFC3339)
	u.Lock = true
	err = us.database.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) UnLock(user *model.User) error {
	u, err := us.database.GetUserByName(user)
	if err != nil {
		return nil
	}
	u.UnlockTime = time.Now().Format(time.RFC3339)
	u.Lock = false
	err = us.database.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

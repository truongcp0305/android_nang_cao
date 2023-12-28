package repository

import (
	"android-service/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-pg/pg/v10"
)

type DatabaseRepo struct {
	db *pg.DB
}

func NewDatabaseRepo(db *pg.DB) *DatabaseRepo {
	return &DatabaseRepo{
		db: db,
	}
}

func (r *DatabaseRepo) GetTaskById(task *model.Task) error {
	err := r.db.Model(task).Where("id = ?", task.Id).First()
	return err
}

func (r *DatabaseRepo) GetListTaskByUserId(task *model.Task) ([]model.Task, error) {
	tasks := []model.Task{}
	err := r.db.Model(task).Where("user_id = ?", task.UserId).Select(&tasks)
	return tasks, err
}

func (r *DatabaseRepo) CreateTask(task *model.Task) error {
	_, err := r.db.Model(task).Insert()
	return err
}

func (r *DatabaseRepo) UpdateTask(task *model.Task) error {
	_, err := r.db.Model(task).WherePK().Update()
	return err
}

func (r *DatabaseRepo) DeleteTask(task *model.Task) error {
	_, err := r.db.Model(task).WherePK().Delete()
	return err
}

func (r *DatabaseRepo) CreateUser(user *model.User) error {
	_, err := r.db.Model(user).Insert()
	return err
}

func (r *DatabaseRepo) GetUserByUserNameAndPass(user *model.User) error {
	err := r.db.Model(user).Where("user_name = ?", user.UserName).Where("password = ?", user.Pass).First()
	return err
}

func (r *DatabaseRepo) GetUserByName(user *model.User) error {
	err := r.db.Model(user).Where("user_name = ?", user.UserName).First()
	return err
}

func (r *DatabaseRepo) UpdateUser(user *model.User) error {
	_, err := r.db.Model(user).Where("id = ?", user.UserId).Update()
	return err
}

func (r *DatabaseRepo) GetUserInfo(info *model.UserInfo) error {
	err := r.db.Model(info).Where("user_id = ?", info.UserId).First()
	return err
}

func (r *DatabaseRepo) CreateUserInfo(info *model.UserInfo) error {
	_, err := r.db.Model(info).Insert()
	return err
}

func (r *DatabaseRepo) UpdateUserInfo(info *model.UserInfo) error {
	_, err := r.db.Model(info).Where("user_id = ?", info.UserId).Update()
	return err
}

func (r *DatabaseRepo) GetListUser() ([]model.User, error) {
	user := model.User{}
	list := []model.User{}
	err := r.db.Model(&user).Where("user_name IS NOT NULL").Select(&list)
	return list, err
}

func (r *DatabaseRepo) GetUserName(user *model.User) error {
	err := r.db.Model(user).Where("user_name = ?", user.UserName).First()
	return err
}

func (r *DatabaseRepo) InsertWord(words []model.Word) error {
	_, err := r.db.Model(&words).Insert()
	return err
}

func (r *DatabaseRepo) GetWordsForQuestion(level string) ([]model.Word, error) {
	result := []model.Word{}
	listQuest := []model.Word{}
	_, err := r.db.Query(&result, "SELECT * FROM word WHERE word.level = ? ORDER BY RANDOM() LIMIT 20", level)
	for _, w := range result {
		if len(listQuest) == 5 {
			return listQuest, nil
		}
		ckeck, audio, err := CheckQuestionAudio(w.Text)
		if err != nil {
			return nil, err
		}
		if ckeck {
			listQuest = append(listQuest, model.Word{
				Text:  w.Text,
				Audio: audio,
				Level: w.Level,
				Id:    w.Id,
			})
		}
	}
	fmt.Println(listQuest)
	return result, err
}

func (r *DatabaseRepo) GetAllTask() ([]model.Task, error) {
	tasks := []model.Task{}
	t := model.Task{}
	err := r.db.Model(&t).Select(&tasks)
	return tasks, err
}

func CheckQuestionAudio(word string) (bool, string, error) {
	res, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word))
	if err != nil {
		return false, "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return false, "", err
	}
	response := []interface{}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, "", err
	}
	for _, item := range response {
		if data, ok := item.(map[string]interface{}); ok {
			if phonetics, ok := data["phonetics"].([]interface{}); ok {
				for _, phonetic := range phonetics {
					if phoneticData, ok := phonetic.(map[string]interface{}); ok {
						if phoneticData["audio"] != "" {
							return true, phoneticData["audio"].(string), nil
						}
					}
				}
			}
		}

	}
	return false, "", nil
}

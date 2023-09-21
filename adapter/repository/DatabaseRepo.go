package repository

import (
	"android-service/model"

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
	_, err := r.db.Model(task).Update()
	return err
}

func (r *DatabaseRepo) DeleteTask(task *model.Task) error {
	_, err := r.db.Model(task).Delete()
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

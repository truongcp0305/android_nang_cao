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

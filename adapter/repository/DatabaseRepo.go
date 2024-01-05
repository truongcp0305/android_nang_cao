package repository

import (
	"android-service/model"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseRepo struct {
	u *mongo.Collection
	t *mongo.Collection
	i *mongo.Collection
}

func NewDatabaseRepo(u *mongo.Collection, t *mongo.Collection, i *mongo.Collection) *DatabaseRepo {
	return &DatabaseRepo{
		u: u,
		t: t,
		i: i,
	}
}

func (r *DatabaseRepo) GetTaskById(task *model.Task) (*model.Task, error) {
	filter := bson.M{"id": task.Id}
	c, err := r.t.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	ts, err := toTask(c)
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, errors.New("not found")
	}
	task = &ts[0]
	return task, nil
	// err := r.db.Model(task).Where("id = ?", task.Id).First()
	// return err
}

func toTask(cursor *mongo.Cursor) ([]model.Task, error) {
	var tasks []model.Task
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var task model.Task
		err := cursor.Decode(&task)
		if err != nil {
			log.Printf("Error decoding task: %v\n", err)
			return nil, errors.New("Internal server error")
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over tasks: %v\n", err)
		return nil, errors.New("Internal server error")
	}

	return tasks, nil
}

func toUser(cursor *mongo.Cursor) ([]model.User, error) {
	var users []model.User
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Printf("Error decoding task: %v\n", err)
			return nil, errors.New("Internal server error")
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over tasks: %v\n", err)
		return nil, errors.New("Internal server error")
	}

	return users, nil
}

func toInfo(cursor *mongo.Cursor) ([]model.UserInfo, error) {
	var useris []model.UserInfo
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var useri model.UserInfo
		err := cursor.Decode(&useri)
		if err != nil {
			log.Printf("Error decoding task: %v\n", err)
			return nil, errors.New("Internal server error")
		}
		useris = append(useris, useri)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Error iterating over tasks: %v\n", err)
		return nil, errors.New("Internal server error")
	}

	return useris, nil
}

func (r *DatabaseRepo) GetListTaskByUserId(task *model.Task) ([]model.Task, error) {
	filter := bson.M{"assign_to": task.UserId}
	c, err := r.t.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return toTask(c)
	// err := r.db.Model(task).Where("assign_to = ?", task.UserId).Select(&tasks)
	// return tasks, err
}

func (r *DatabaseRepo) CreateTask(task *model.Task) error {
	_, err := r.t.InsertOne(context.Background(), task)
	return err
}

func (r *DatabaseRepo) UpdateTask(task *model.Task) error {
	filter := bson.M{"id": task.Id} // Assuming the task ID is stored in the "_id" field

	update := bson.M{"$set": task}

	_, err := r.t.UpdateOne(context.TODO(), filter, update)
	return err
	// _, err := r.db.Model(task).WherePK().Update()
	// return err
}

func (r *DatabaseRepo) DeleteTask(task *model.Task) error {
	filter := bson.M{"id": task.Id}
	_, err := r.t.DeleteOne(context.Background(), filter)
	return err
	// _, err := r.db.Model(task).WherePK().Delete()
	// return err
}

func (r *DatabaseRepo) CreateUser(user *model.User) error {
	_, err := r.u.InsertOne(context.Background(), user)
	return err
}

func (r *DatabaseRepo) GetUserByUserNameAndPass(user *model.User) (*model.User, error) {
	filter := bson.M{"user_name": user.UserName, "password": user.Pass}
	c, err := r.u.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	us, err := toUser(c)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, errors.New("not found")
	}
	return &us[0], nil
	// err := r.db.Model(user).Where("user_name = ?", user.UserName).Where("password = ?", user.Pass).First()
	// return err
}

func (r *DatabaseRepo) GetUserByName(user *model.User) (*model.User, error) {
	filter := bson.M{"user_name": user.UserName}
	c, err := r.u.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	us, err := toUser(c)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, errors.New("not found")
	}
	return &us[0], nil
	// err := r.db.Model(user).Where("user_name = ?", user.UserName).First()
	// return err
}

func (r *DatabaseRepo) UpdateUser(user *model.User) error {
	filter := bson.M{"id": user.UserId} // Assuming the task ID is stored in the "_id" field
	update := bson.M{"$set": user}
	_, err := r.u.UpdateOne(context.TODO(), filter, update)
	return err
	// _, err := r.db.Model(user).Where("id = ?", user.UserId).Update()
	// return err
}

func (r *DatabaseRepo) GetUserInfo(info *model.UserInfo) (*model.UserInfo, error) {
	filter := bson.M{"user_id": info.UserId}
	c, err := r.i.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	us, err := toInfo(c)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, errors.New("not found")
	}
	info = &us[0]
	return info, nil
	// err := r.db.Model(info).Where("user_id = ?", info.UserId).First()
	// return err
}

func (r *DatabaseRepo) GetUserTasks(user model.User) (int, error) {
	f := bson.M{"user_id": user.UserId, "status": bson.D{{Key: "$ne", Value: "CLOSE"}}}
	c, err := r.t.CountDocuments(context.Background(), f)
	return int(c), err
	// task := model.Task{
	// 	UserId: user.UserId,
	// }
	// c, err := r.db.Model(&task).Where("user_id = ?", user.UserId).Where("status != 'CLOSE'").Count()
	// return c, err
}

func (r *DatabaseRepo) GetDoneTasks(user model.User) (int, error) {
	f := bson.M{"user_id": user.UserId, "status": "DONE"}
	c, err := r.t.CountDocuments(context.Background(), f)
	return int(c), err
	// task := model.Task{
	// 	UserId: user.UserId,
	// }
	// c, err := r.db.Model(&task).Where("user_id = ?", user.UserId).Where("status = 'DONE'").Count()
	// return c, err
}

func (r *DatabaseRepo) GetAssignTasks(task model.Task) ([]model.Task, error) {
	filter := bson.M{"user_id": task.UserId, "assign_to": bson.M{"$ne": task.UserId}}
	c, err := r.t.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return toTask(c)
	// tasks := []model.Task{}
	// err := r.db.Model(&task).Where("user_id = ?", task.UserId).Where("assign_to != ?", task.UserId).Select(&tasks)
	// return tasks, err
}

func (r *DatabaseRepo) CreateUserInfo(info *model.UserInfo) error {
	_, err := r.i.InsertOne(context.Background(), info)
	return err
	// _, err := r.db.Model(info).Insert()
	// return err
}

func (r *DatabaseRepo) UpdateUserInfo(info *model.UserInfo) error {
	f := bson.M{"user_id": info.UserId}
	n := bson.M{"$set": info}
	_, err := r.i.UpdateOne(context.Background(), f, n)
	return err
	// _, err := r.db.Model(info).Where("user_id = ?", info.UserId).Update()
	// return err
}

func (r *DatabaseRepo) GetListUser() ([]model.User, error) {
	filter := bson.M{}
	c, err := r.u.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	return toUser(c)
	// user := model.User{}
	// list := []model.User{}
	// err := r.db.Model(&user).Where("user_name IS NOT NULL").Select(&list)
	// return list, err
}

func (r *DatabaseRepo) GetUserName(user *model.User) (*model.User, error) {
	filter := bson.M{"user_name": user.UserName}
	c, err := r.u.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	us, err := toUser(c)
	if err != nil {
		return nil, err
	}
	if len(us) == 0 {
		return nil, errors.New("not found")
	}
	user = &us[0]
	return user, nil
	// err := r.db.Model(user).Where("user_name = ?", user.UserName).First()
	// return err
}

func (r *DatabaseRepo) GetAllTask() ([]model.Task, error) {
	f := bson.M{}
	c, err := r.t.Find(context.Background(), f)
	if err != nil {
		return nil, err
	}
	return toTask(c)
	// tasks := []model.Task{}
	// t := model.Task{}
	// err := r.db.Model(&t).Select(&tasks)
	// return tasks, err
}

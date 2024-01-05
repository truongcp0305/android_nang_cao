package registry

import (
	"android-service/adapter/controller"
	"android-service/adapter/repository"
	"android-service/usecase/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Registry struct {
	db *mongo.Database
}

func NewRegistry(db *mongo.Database) Registry {
	return Registry{
		db: db,
	}
}

func (r *Registry) NewTaskController() controller.TaskController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewTaskController(
		*service.NewTaskService(db),
	)
}

func (r *Registry) NewUserController() controller.UserController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewUserController(
		*service.NewUserService(db),
	)
}

func (r *Registry) NewWordController() controller.WordController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewWordController(
		*service.NewWordService(db),
	)
}

func (r *Registry) NewSocketController() controller.SocketController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewSocketController(
		*service.NewSocketService(db),
		*service.NewRoom(),
	)
}

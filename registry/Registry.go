package registry

import (
	"android-service/adapter/controller"
	"android-service/adapter/repository"
	"android-service/usecase/service"

	"github.com/go-pg/pg/v10"
)

type Registry struct {
	db *pg.DB
}

func NewRegistry(db *pg.DB) Registry {
	return Registry{
		db: db,
	}
}

func (r *Registry) NewTaskController() controller.TaskController {
	db := repository.NewDatabaseRepo(r.db)
	return controller.NewTaskController(
		*service.NewTaskService(db),
	)
}

func (r *Registry) NewUserController() controller.UserController {
	db := repository.NewDatabaseRepo(r.db)
	return controller.NewUserController(
		*service.NewUserService(db),
	)
}

func (r *Registry) NewWordController() controller.WordController {
	db := repository.NewDatabaseRepo(r.db)
	return controller.NewWordController(
		*service.NewWordService(db),
	)
}

func (r *Registry) NewSocketController() controller.SocketController {
	db := repository.NewDatabaseRepo(r.db)
	return controller.NewSocketController(
		*service.NewSocketService(db),
		*service.NewRoom(),
	)
}

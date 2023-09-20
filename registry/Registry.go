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

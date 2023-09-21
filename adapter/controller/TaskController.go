package controller

import (
	"android-service/adapter/incoming"
	"android-service/usecase/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	task service.TaskService
}

func NewTaskController(task service.TaskService) TaskController {
	return TaskController{
		task: task,
	}
}

// @Summary Detail a task
// @Tags Task
// @Param id path string true "task id" example(1101)
// @Success 200 {object} model.Task
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /task/{id} [get]
func (Cs *TaskController) Detail(c echo.Context) error {
	var params incoming.DetailTaskParams
	c.Bind(&params)
	if params.Id == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	task := params.GetModel()
	err := Cs.task.DetailTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": task,
	})
}

// @Summary Get list task
// @Tags Task
// @Param userId path string true "userId"
// @Success 200 {array} model.Task
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /task/list-task/{userId} [get]
func (Cs *TaskController) GetList(c echo.Context) error {
	var params incoming.GetListTaskParam
	c.Bind(&params)
	if params.UserId == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	task := params.GetModel()
	tasks, err := Cs.task.GetList(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": tasks,
	})
}

// @Summary Create a task
// @Tags Task
// @Param data body string true "json of a task"
// @Success 200 {object} outgoing.ModelReturn
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /task [post]
func (Cs *TaskController) Create(c echo.Context) error {
	var params incoming.CreateTaskParams
	c.Bind(&params)
	if params.Data == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.task.CreateTask(params.Data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Create task sucsess",
	})
}

// @Summary Update a task
// @Tags Task
// @Param data body string true "json of a task"
// @Success 200 {object} outgoing.ModelReturn
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /task [put]
func (Cs *TaskController) Update(c echo.Context) error {
	var params incoming.UpdateTaskParams
	c.Bind(&params)
	if params.Data == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.task.UpdateTask(params.Data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Update sucsess",
	})
}

// @Summary Delete a task
// @Tags Task
// @Param id path string true "task id" example(1101)
// @Success 200 {object} outgoing.ModelReturn
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /task/{id} [delete]
func (Cs *TaskController) Delete(c echo.Context) error {
	var params incoming.DeleteTaskParams
	c.Bind(&params)
	if params.Id == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	task := params.GetModel()
	err := Cs.task.DeleteTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Delete sucsess",
	})
}

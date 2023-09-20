package incoming

import "android-service/model"

type DetailTaskParams struct {
	Id string `json:"id" param:"id" query:"id" form:"id"`
}

func (ic *DetailTaskParams) GetModel() *model.Task {
	return &model.Task{
		Id: ic.Id,
	}
}

type GetListTaskParam struct {
	UserId string `json:"userId" param:"userId" query:"userId" form:"userId"`
}

func (ic *GetListTaskParam) GetModel() *model.Task {
	return &model.Task{
		UserId: ic.UserId,
	}
}

type CreateTaskParams struct {
	Data string `json:"data" form:"data"`
}

type UpdateTaskParams struct {
	Data string `json:"data" form:"data"`
}

type DeleteTaskParams struct {
	Id string `json:"id" query:"id" form:"id"`
}

func (ic *DeleteTaskParams) GetModel() *model.Task {
	return &model.Task{
		Id: ic.Id,
	}
}

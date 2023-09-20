package model

type Task struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	UserId     string   `json:"userId" pg:"user_id"`
	Desciption string   `json:"desciption"`
	Priority   string   `json:"priority"`
	TaskType   string   `json:"type"`
	AssignTo   string   `json:"assignTo" pg:"assign_to"`
	AssignName string   `json:"assignName"`
	Sprint     string   `json:"sprint"`
	Status     string   `json:"status"`
	CreateTime string   `json:"createTime" pg:"create_time"`
	StartTime  string   `json:"startTime" pg:"start_time"`
	EndTime    string   `json:"endTime" pg:"end_time"`
	Comment    string   `json:"comment"`
	tableName  struct{} `pg:"task"`
}

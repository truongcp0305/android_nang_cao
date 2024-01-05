package model

type Task struct {
	Id         string `json:"id" bson:"id"`
	Name       string `json:"name" bson:"name"`
	UserId     string `json:"userId" bson:"user_id"`
	Desciption string `json:"desciption" bson:"desciption"`
	Priority   string `json:"priority" bson:"priority"`
	TaskType   string `json:"taskType" bson:"task_type"`
	AssignTo   string `json:"assignTo" bson:"assign_to"`
	AssignName string `json:"assignName" bson:"assign_name"`
	Sprint     string `json:"sprint" bson:"sprint"`
	Status     string `json:"status" bson:"status"`
	CreateTime string `json:"createTime" bson:"create_time"`
	StartTime  string `json:"startTime" bson:"start_time"`
	EndTime    string `json:"endTime" bson:"end_time"`
	Comment    string `json:"comment" bson:"comment"`
	//tableName  struct{} `pg:"task"`
}

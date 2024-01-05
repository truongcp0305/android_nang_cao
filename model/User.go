package model

type User struct {
	UserId        string `json:"userId" bson:"id"`
	Lock          bool   `json:"lock" bson:"lock"`
	UserName      string `json:"userName" bson:"user_name"`
	Role          string `json:"role" bson:"role"`
	Pass          string `json:"pass" bson:"password"`
	Try           int    `json:"try" bson:"try"`
	UnlockTime    string `json:"unlockTime" bson:"unlock_time"`
	Token         string `json:"token" bson:"-"`
	TotalTask     int    `json:"totalTask" bson:"total_task"`
	TotalTaskDone int    `josn:"totalTaskDone" bson:"total_task_done"`
	//tableName     struct{} `bson:"user"`
}

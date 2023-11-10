package model

type UserInfo struct {
	Id         string   `json:"id"`
	UserId     string   `json:"userId" pg:"user_id"`
	UserName   string   `json:"userName" pg:"user_name"`
	Point      string   `json:"point"`
	OtherInfor string   `json:"otherInfor" pg:"other_infor"`
	tableName  struct{} `pg:"user_infor"`
}

package model

type UserInfo struct {
	Id         string `json:"id"`
	UserId     string `json:"userId" bson:"user_id"`
	UserName   string `json:"userName" bson:"user_name"`
	Point      string `json:"point" bson:"point"`
	OtherInfor string `json:"otherInfor" bson:"other_infor"`
	Role       string `json:"role" bson:"role"`
	Token      string `json:"token" bson:"-"`
	//tableName  struct{} `pg:"user_infor"`
}

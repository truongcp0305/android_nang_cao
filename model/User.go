package model

type User struct {
	UserId    string   `json:"userId" pg:"id"`
	UserName  string   `json:"userName" pg:"user_name"`
	Pass      string   `json:"pass"`
	tableName struct{} `pg:"user"`
}
package model

type User struct {
	UserId     string   `json:"userId" pg:"id"`
	UserName   string   `json:"userName" pg:"user_name"`
	Role       string   `json:"role" pg:"role"`
	Pass       string   `json:"pass" pg:"password"`
	Try        int      `json:"try" pg:"try"`
	UnlockTime string   `json:"unlockTime" pg:"unlock_time"`
	Token      string   `json:"token" pg:"-"`
	tableName  struct{} `pg:"user"`
}

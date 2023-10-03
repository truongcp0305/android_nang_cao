package model

type Word struct {
	Id        string   `josn:"id"`
	Text      string   `json:"text"`
	Level     int64    `json:"level"`
	tableName struct{} `pg:"word"`
}

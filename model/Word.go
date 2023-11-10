package model

type Word struct {
	Id        string   `josn:"id"`
	Text      string   `json:"text"`
	Level     int64    `json:"level"`
	Audio     string   `json:"audio"`
	tableName struct{} `pg:"word"`
}

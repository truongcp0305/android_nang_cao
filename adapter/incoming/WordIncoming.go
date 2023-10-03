package incoming

type InsertWord struct {
	Texts string `json:"texts" form:"texts"`
	Level string `json:"level" form:"level"`
}

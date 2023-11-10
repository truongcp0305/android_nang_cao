package incoming

type InsertWord struct {
	Texts string `json:"texts" form:"texts"`
	Level string `json:"level" form:"level"`
}

type GetQuestion struct {
	Level string `json:"level" query:"level" param:"level"`
}

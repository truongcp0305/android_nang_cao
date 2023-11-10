package model

type MatchStatus struct {
	Id        string `json:"Id"`
	Message   string `json:"message"`
	Status    string `json:"status"`
	Point     string `json:"point"`
	Questions []Word `json:"questions"`
}

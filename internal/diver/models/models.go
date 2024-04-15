package models

type Report struct {
	Id    string `json:"id"`
	Query string `json:"query"`
}

type Treasure struct {
	Id   string `json:"id"`
	Data string `json:"data"`
}

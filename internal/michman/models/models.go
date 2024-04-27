package models

type Diver struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Addr    string   `json:"addr"`
	Reports []Report `json:"reports,omitempty"`
}

type Report struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Query string `json:"query"`
}

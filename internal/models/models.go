package models

type Report struct {
	Id    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Query string `json:"query,omitempty"`
}

type Data struct {
	ID      string
	Columns []string `json:"columns"`
	Values  [][]any  `json:"values"`
}

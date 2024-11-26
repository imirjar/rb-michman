package models

type Report struct {
	Id   string `json:"id"`
	Name string `json:"name,omitempty"`
	Data string `json:"data,omitempty"`
}

type Target struct {
	DiverId  string `json:"diver_id"`
	ReportId string `json:"report_id"`
}

type Diver struct {
	Name      string   `json:"name"`
	Addr      string   `json:"addr"`
	Reports   []Report `json:"reports,omitempty"`
	Connected bool
}

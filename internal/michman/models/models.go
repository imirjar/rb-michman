package models

type Diver struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Addr    string   `json:"addr"`
	Reports []Report `json:"reports,omitempty"`
}

type Report struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Data string `json:"data,omitempty"`
}

// type ExecReport struct {
// 	DiverId  string `json:"diver_id"`
// 	ReportId string `json:"report_id"`
// 	Report   Report `json:"report"`
// }

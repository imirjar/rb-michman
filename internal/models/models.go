package models

import (
	"net/http"
	"time"
)

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
	Name    string   `json:"name"`
	Addr    string   `json:"addr"`
	Reports []Report `json:"reports,omitempty"`
}

func (d *Diver) CheckConn() bool {
	clinet := http.Client{
		Timeout: 3 * time.Second,
	}
	response, err := clinet.Get("http://" + d.Addr + "/check_conn")
	if err != nil {
		return false
	}
	defer response.Body.Close()

	return response.StatusCode == 200
}

type Data struct {
	Columns []string `json:"columns"`
	Values  [][]any  `json:"values"`
}

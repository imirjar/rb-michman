package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/imirjar/Michman/internal/michman/models"
)

type API struct {
	Client      http.Client
	ContentType string
}

func NewAPI() *API {
	api := API{
		Client: http.Client{
			Timeout: 3 * time.Second,
		},
		ContentType: "application/json",
	}
	return &api
}

func (api API) GetDiverReports(ctx context.Context, path string) ([]models.Report, error) {
	var reports []models.Report
	response, err := api.Client.Post(path+"/reports/list/", api.ContentType, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		return reports, errors.New(response.Status)
	}
	err = json.NewDecoder(response.Body).Decode(&reports)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (api API) ExecuteDiverReport(ctx context.Context, addr, repId string) (models.Report, error) {
	var report = models.Report{
		Id: repId,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(report)
	if err != nil {
		return report, err
	}

	response, err := api.Client.Post(addr+"/reports/execute/", api.ContentType, &buf)
	if err != nil {
		return report, err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		return report, errors.New(response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(&report)
	if err != nil {
		return report, err
	}
	return report, nil
}

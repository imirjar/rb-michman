package diver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/imirjar/rb-michman/internal/models"
)

type API struct {
	Client      http.Client
	ContentType string
}

func New() *API {
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
	log.Print(path)
	response, err := api.Client.Get("http://" + path + "/reports/")
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

func (api API) ExecuteDiverReport(ctx context.Context, addr, repId string) (models.Data, error) {
	report := models.Data{}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(report)
	if err != nil {
		return report, err
	}

	response, err := api.Client.Get("http://" + addr + "/reports/generate/" + repId)
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

func (api API) CheckConnection(ctx context.Context, dvr models.Diver) bool {
	return true
}

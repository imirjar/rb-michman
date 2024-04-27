package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/imirjar/Michman/internal/michman/models"
)

type API struct {
	Client      http.Client
	Path        string
	ContentType string
}

func NewAPI() *API {
	api := &API{
		Client:      http.Client{},
		ContentType: "application/json",
	}
	return api
}

func (api API) GetDiverReports(ctx context.Context, path string) ([]models.Report, error) {
	var reports []models.Report
	response, err := api.Client.Post(path, api.ContentType, nil)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&reports)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

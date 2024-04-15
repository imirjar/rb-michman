package api

import (
	"context"
	"io"
	"net/http"
)

type API struct {
	Client http.Client
	Path   string
}

func NewAPI() *API {
	api := &API{
		Client: http.Client{},
		Path:   "http://localhost:8080/reports/list/",
	}
	return api
}

func (api API) GetDiverReports(ctx context.Context, id string) (string, error) {

	response, err := api.Client.Post(api.Path, "application/json", nil)
	if err != nil {
		return err.Error(), err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body) // response body is []byte
	if err != nil {
		return err.Error(), err
	}

	return string(body), nil
}

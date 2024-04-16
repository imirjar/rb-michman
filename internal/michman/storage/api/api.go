package api

import (
	"context"
	"io"
	"net/http"
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

func (api API) GetDiverReports(ctx context.Context, path string) (string, error) {

	response, err := api.Client.Post(path, api.ContentType, nil)
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

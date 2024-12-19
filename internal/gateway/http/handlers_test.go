package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/imirjar/rb-michman/internal/mock"
	"github.com/imirjar/rb-michman/internal/models"
)

type (
	diverService struct {
		reports []models.Report
		data    models.Data
		err     error
	}

	grazerService struct {
		diversMap map[string]models.Diver
		diver     models.Diver
		err       error
	}

	request struct {
		url    string
		format string
	}

	response struct {
		contentType string
		code        int
		body        string
	}

	mockService struct {
		ctrl   *gomock.Controller
		diver  *mock.MockDiver
		grazer *mock.MockGrazer
	}

	testApp struct {
		App
	}
)

func newTestApp(s *mockService) *App {
	app := App{
		GrazerService: s.grazer,
		DiverService:  s.diver,
	}
	return &app
}

func (s *mockService) withGrazer() *mockService {
	s.grazer = mock.NewMockGrazer(s.ctrl)
	return s
}

func (s *mockService) withDiver() *mockService {
	s.diver = mock.NewMockDiver(s.ctrl)
	return s
}

func newMockService(t *testing.T) *mockService {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return &mockService{
		ctrl:   ctrl,
		diver:  mock.NewMockDiver(ctrl),
		grazer: mock.NewMockGrazer(ctrl),
	}
}

func NewTestApp(t *testing.T) *App {

	return &App{}
}

func TestInfo(t *testing.T) {

	tests := []struct {
		name       string
		mockReturn map[string]models.Diver
		resp       response
	}{
		{
			name: "ok",
			mockReturn: map[string]models.Diver{
				"1": models.Diver{
					Name: "one",
					Addr: "localhost:9091",
				},
				"2": models.Diver{
					Name: "two",
					Addr: "localhost:9092",
				},
			},
			resp: response{
				contentType: "text/html",
				code:        http.StatusOK,
			},
		},
		{
			name:       "no divers",
			mockReturn: map[string]models.Diver{},
			resp: response{
				contentType: "text/html",
				code:        http.StatusOK,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := New()
	grazer := mock.NewMockGrazer(ctrl)
	app.GrazerService = grazer

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			grazer.EXPECT().DiverList(gomock.Any()).Return(tt.mockReturn, nil)

			server := httptest.NewServer(http.HandlerFunc(app.Info()))
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != tt.resp.code {
				t.Errorf("CODE %d, expected %d", resp.StatusCode, tt.resp.code)
			}

		})

	}
}

func TestGetDivers(t *testing.T) {

	tests := []struct {
		name       string
		mockReturn map[string]models.Diver
		resp       response
	}{
		{
			name: "ok",
			mockReturn: map[string]models.Diver{
				"1": models.Diver{
					Name: "one",
					Addr: "localhost:9091",
				},
				"2": models.Diver{
					Name: "two",
					Addr: "localhost:9092",
				},
			},
			resp: response{
				contentType: "",
				code:        200,
			},
		},
		{
			name:       "no divers",
			mockReturn: map[string]models.Diver{},
			resp: response{
				contentType: "application/json",
				code:        http.StatusOK,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := New()
	diver := mock.NewMockDiver(ctrl)
	grazer := mock.NewMockGrazer(ctrl)
	app.DiverService = diver
	app.GrazerService = grazer

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			grazer.EXPECT().DiverList(gomock.Any()).Return(tt.mockReturn, nil)

			server := httptest.NewServer(http.HandlerFunc(app.GetDivers()))
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != tt.resp.code {
				t.Errorf("CODE %d, expected %d", resp.StatusCode, tt.resp.code)
			}

		})

	}

}

func TestGetDiverReports(t *testing.T) {

	tests := []struct {
		name    string
		service diverService
		resp    response
	}{
		{
			name: "ok",
			service: diverService{
				reports: []models.Report{
					models.Report{
						Id:   "1",
						Name: "one",
						Data: "",
					},
				},
				err: nil,
			},
			resp: response{
				contentType: "application/json",
				code:        200,
			},
		},
		{
			name: "no reports",
			service: diverService{
				reports: []models.Report{},
				err:     nil,
			},
			resp: response{
				contentType: "application/json",
				code:        200,
			},
		},
		{
			name: "with service error",
			service: diverService{
				reports: []models.Report{},
				err:     errors.New("some error"),
			},
			resp: response{
				contentType: "application/json",
				code:        500,
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := New()
	diver := mock.NewMockDiver(ctrl)
	app.DiverService = diver

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			diver.EXPECT().DiverReports(gomock.Any(), "").Return(tt.service.reports, tt.service.err)

			server := httptest.NewServer(http.HandlerFunc(app.GetDiverReports()))
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Error(err)
			}

			if resp.StatusCode != tt.resp.code {
				t.Errorf("CODE %d, expected %d", resp.StatusCode, tt.resp.code)
			}

			ct := resp.Header.Get("Content-type")
			if ct != tt.resp.contentType {
				t.Errorf("Content type is %s, expected %s", ct, tt.resp.contentType)
			}

		})

	}

}

func TestExecuteDiverReport(t *testing.T) {
	tests := []struct {
		name string

		request  request
		response response

		grazerService grazerService
		diverService  diverService
	}{
		{
			name: "ok",
			request: request{
				url:    "/1/1",
				format: "",
			},
			response: response{
				contentType: "application/json",
				code:        http.StatusOK,
			},
			diverService: diverService{
				data: models.Data{
					Columns: []string{"column1", "column2"},
					Values:  [][]any{{"value1", "value2"}},
				},
				err: nil,
			},
			grazerService: grazerService{
				diversMap: map[string]models.Diver{
					"1": models.Diver{},
				},
				err: nil,
			},
		},
		{
			name: "service error",
			request: request{
				url:    "/1/1",
				format: "",
			},
			diverService: diverService{
				data: models.Data{},
				err:  errors.New("service error"),
			},
			grazerService: grazerService{
				diversMap: map[string]models.Diver{
					"1": models.Diver{},
				},
				err: nil,
			},
			response: response{
				contentType: "application/json",
				code:        http.StatusBadRequest,
			},
		},
	}

	service := newMockService(t).withDiver().withGrazer()
	app := newTestApp(service)

	router := chi.NewRouter()
	router.Get("/{id}/{reportId}", app.ExecuteDiverReport())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			service.grazer.EXPECT().DiverAddr(gomock.Any(), gomock.Any()).Return("validAddress", tt.grazerService.err)
			if tt.grazerService.err != nil {
				return
			}
			service.diver.EXPECT().GetDiverReportData(gomock.Any(), "validAddress", gomock.Any()).Return(tt.diverService.data, tt.diverService.err)

			server := httptest.NewServer(router)
			defer server.Close()

			reqURL := server.URL + tt.request.url
			// if tt.request.format != "" {
			// 	reqURL += "?format=" + tt.queryFormat
			// }

			resp, err := http.Get(reqURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.response.code {
				t.Errorf("expected status %d, got %d", tt.response.code, resp.StatusCode)
			}

			if resp.Header.Get("Content-type") != tt.response.contentType {
				t.Errorf("expected content-type %s, got %s", tt.response.contentType, resp.Header.Get("Content-type"))
			}
		})
	}
}

package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/app/middleware"
	"github.com/imirjar/Michman/internal/michman/models"
	"github.com/imirjar/Michman/internal/michman/service"
)

type Service interface {
	DiverList(context.Context) ([]models.Diver, error)
	DiverReports(context.Context, string) ([]models.Report, error)
	GetDiverReportData(ctx context.Context, ex models.Diver) (models.Report, error)
}

type App struct {
	config  Config
	Service Service
}

type Config interface {
	GetAuthAddr() string
	GetMichmanAddr() string //allow req only for this addr
	GetSecret() string
}

func NewApp() *App {

	return &App{
		config:  config.NewConfig(),
		Service: service.NewService(),
	}
}

func (a *App) Run(ctx context.Context) error {
	router := chi.NewRouter()

	router.Use(middleware.Encryptor(a.config.GetSecret(), a.config.GetAuthAddr()))
	router.Use(middleware.REST())

	router.Get("/", a.Hello)

	router.Post("/divers/", a.DiversListHandler())

	router.Route("/diver", func(diver chi.Router) {
		diver.Post("/reports/", a.DiverReportsHandler())
		diver.Post("/execute/", a.DiverExecuteReportHandler())
	})

	srv := &http.Server{
		Addr:    a.config.GetMichmanAddr(),
		Handler: router,
	}

	log.Printf("Run app on PORT %s", a.config.GetMichmanAddr())
	return srv.ListenAndServe()
}

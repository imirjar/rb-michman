package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/diver/service"
)

type Service interface {
	Execute(ctx context.Context, id string) (string, error)
	ReportsList(ctx context.Context) (string, error)
}

type App struct {
	config  Config
	Secret  string
	Service Service
	Server  *http.Server
}

type Config interface {
	GetDiverAddr() string
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
	router.Route("/reports", func(update chi.Router) {
		update.Post("/execute/", a.ExecuteHandler)
		update.Post("/list/", a.ReportsListHandler)
	})

	//for new usecases add new route
	srv := &http.Server{
		Addr:    a.config.GetDiverAddr(),
		Handler: router,
	}

	log.Printf("Run app on PORT %s", a.config.GetDiverAddr())
	return srv.ListenAndServe()
}

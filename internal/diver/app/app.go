package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/internal/diver/service"
)

type Service interface {
	Execute(ctx context.Context, id string) (string, error)
	ReportsList(ctx context.Context) (string, error)
}

type App struct {
	Service Service
}

func NewApp(addr string) *App {
	return &App{
		Service: service.NewService(),
	}
}

func (a *App) Run(ctx context.Context) error {
	log.Print("Run app")

	router := chi.NewRouter()
	router.Route("/reports", func(update chi.Router) {
		update.Post("/execute/", a.ExecuteHandler)
		update.Post("/list/", a.ReportsListHandler)
	})

	//for new usecases add new route

	gateway := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	return gateway.ListenAndServe()
}

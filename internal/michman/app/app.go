package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/internal/michman/service"
)

type Service interface {
	DiverList(context.Context) (string, error)
	DiverInfo(context.Context, string) (string, error)
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

	router.Get("/", a.Hello)

	router.Post("/divers/", a.DiversListHandler())

	router.Route("/diver", func(diver chi.Router) {
		diver.Post("/info/", a.DiverInfoHandler())
	})

	gateway := &http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	return gateway.ListenAndServe()
}

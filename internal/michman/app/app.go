package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/service"
)

type Service interface {
	DiverList(context.Context) (string, error)
	DiverInfo(context.Context, string) (string, error)
}

type App struct {
	config  Config
	Service Service
}

type Config interface {
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

	router.Get("/", a.Hello)

	router.Post("/divers/", a.DiversListHandler())

	router.Route("/diver", func(diver chi.Router) {
		diver.Post("/info/", a.DiverInfoHandler())
	})

	srv := &http.Server{
		Addr:    a.config.GetMichmanAddr(),
		Handler: router,
	}

	log.Printf("Run app on PORT %s", a.config.GetMichmanAddr())
	return srv.ListenAndServe()
}

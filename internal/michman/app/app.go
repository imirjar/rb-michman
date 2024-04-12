package app

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/internal/diver/service"
)

type Service interface {
}

type App struct {
	Service Service
}

func NewApp(addr string) *App {
	return &App{
		Service: service.NewService(),
	}
}

// must check whitch of the divers the user have access
func (a *App) DiverList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am Michman"))
}

func (a *App) Run(ctx context.Context) error {
	log.Print("Run app")
	router := chi.NewRouter()
	router.Get("/", a.DiverList)

	gateway := &http.Server{
		Addr:    "localhost:9090",
		Handler: router,
	}

	return gateway.ListenAndServe()
}

package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/Michman/internal/diver/models"
	"github.com/imirjar/Michman/internal/diver/service"
)

type Service interface {
	Execute(ctx context.Context, id string) (string, error)
}

type App struct {
	Service Service
}

func NewApp(addr string) *App {
	return &App{
		Service: service.NewService(),
	}
}

func (a *App) ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var report models.Treasure
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		log.Println("HANDLER ExecuteHandler Decode ERROR", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.Service.Execute(context.Background(), report.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	report.Data = result
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(report); err != nil {
		log.Println("HANDLER ExecuteHandler Encode ERROR", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// w.Write([]byte(report))

}

func (a *App) Run(ctx context.Context) error {
	log.Print("Run app")
	router := chi.NewRouter()
	router.Post("/execute/", a.ExecuteHandler)

	gateway := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	return gateway.ListenAndServe()
}

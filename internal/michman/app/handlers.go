package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/imirjar/Michman/internal/michman/models"
)

// must check whitch of the divers the user have access
func (a *App) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am Michman"))
}

// must check whitch of the divers the user have access
func (a *App) DiversListHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		divers, err := a.Service.DiverList(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(divers))
	}
}

func (a *App) DiverInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// diverId := chi.URLParam(r, "id")
		var diver models.Diver
		err := json.NewDecoder(r.Body).Decode(&diver)
		if err != nil {
			log.Println("HANDLER ExecuteHandler Decode ERROR", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Print(diver)
		divers, err := a.Service.DiverInfo(r.Context(), diver.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(divers))
	}
}

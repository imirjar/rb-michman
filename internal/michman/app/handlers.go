package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/imirjar/Michman/internal/michman/models"
)

// must check whitch of the divers the user have access
func (a *App) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("VSYO OK")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("I am Michman"))
	}
}

func (a *App) DiversList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		divers, err := a.Service.DiverList(r.Context())
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(divers); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (a *App) ReportsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diver models.Diver
		err := json.NewDecoder(r.Body).Decode(&diver)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reports, err := a.Service.DiverReports(r.Context(), diver.Id)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(reports); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (a *App) ReportExecute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diver models.Diver
		err := json.NewDecoder(r.Body).Decode(&diver)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := a.Service.GetDiverReportData(r.Context(), diver)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

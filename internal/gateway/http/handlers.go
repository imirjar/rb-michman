package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-michman/internal/models"
)

// must check whitch of the divers the user have access
func (a *App) Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := `Hi! My name is Michman and I can manage all of your databases in a single API.`
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}
}

func (a *App) DiversList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		divers, err := a.GrazerService.DiverList(r.Context())
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

func (a *App) DiverReportsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r, "id")

		reports, err := a.ReportService.DiverReports(r.Context(), hash)
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

func (a *App) DiverReportExecute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var diver models.Diver
		err := json.NewDecoder(r.Body).Decode(&diver)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := a.ReportService.GetDiverReportData(r.Context(), diver)
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

// Is used for connecting runing divers into Michman grazer
func (a *App) Connect() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Print("connect")

		var diver models.Diver
		diver.Addr = "127.0.0.1:8080"
		diver.Name = "diver"
		// err := json.NewDecoder(r.Body).Decode(&diver)
		// if err != nil {
		// 	log.Print(err)
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }

		err := a.GrazerService.ConnectDiver(r.Context(), diver)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("connected"))
	}
}

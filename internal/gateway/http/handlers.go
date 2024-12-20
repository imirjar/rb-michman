package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-michman/internal/models"
)

// General information about the application
// Also grant a information of connected divers
func (a *App) Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := `Hi! My name is Michman and I can manage all of your databases in a single API.`

		divers, err := a.GrazerService.DiverList(r.Context())
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(divers) > 0 {
			text += `<table style="border-style: solid;"><tr><th>name</th><th>addres</th><th>status</th></tr>`
			for _, v := range divers {
				status := "<a>&#128993;</a>"
				if v.CheckConn() {
					status = "<a>&#128994;</a>"
				} else {
					status = "<a>&#128308;</a>"
				}
				text += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", v.Name, v.Addr, status)
			}
			text += "</table>"
		}

		w.Header().Add("content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}
}

// Get all divers from inmemory storage
func (a *App) GetDivers() http.HandlerFunc {
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

// Connect to choosen diver and get the diver list
func (a *App) GetDiverReports() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		hash := chi.URLParam(r, "id")

		reports, err := a.DiverService.DiverReports(r.Context(), hash)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(reports); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
			return
		}
	}
}

// Connect to choosen diver and execute choosen report. Then get it's data
func (a *App) ExecuteDiverReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")

		hash := chi.URLParam(r, "id")
		repID := chi.URLParam(r, "reportId")

		diverAddr, err := a.GrazerService.DiverAddr(r.Context(), hash)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "status bad request"})
			return
		}

		query := r.FormValue("format")
		if query == "json" {
			data, err := a.DiverService.GetDiverReportDataMap(r.Context(), diverAddr, repID)
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
		} else {
			data, err := a.DiverService.GetDiverReportData(r.Context(), diverAddr, repID)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "status bad request"})
				return
			}
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(data); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": "status bad request"})
				return
			}
		}

	}
}

// Is used for connecting runing divers into Michman inmemory storage
func (a *App) ConnectDiver() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var diver models.Diver

		if err := json.NewDecoder(r.Body).Decode(&diver); err != nil {
			http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
			log.Printf("Ошибка декодирования JSON: %v", err)
			return
		}

		log.Println("DIVER->", diver)

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

func errorJSON(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
}

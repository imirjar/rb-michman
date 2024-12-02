package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-michman/internal/models"
)

// must check whitch of the divers the user have access
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
			text += `<table style="border-style: solid;"><tr><th>name</th><th>addres</th></tr>`
			for _, v := range divers {
				text += fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>", v.Name, v.Addr)
			}
			text += "</table>"
		}

		w.Header().Add("content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}
}

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

func (a *App) GetDiverReports() http.HandlerFunc {
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

func (a *App) ExecuteDiverReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hash := chi.URLParam(r, "id")
		repID := chi.URLParam(r, "reportId")

		diverAddr, err := a.GrazerService.DiverAddr(r.Context(), hash)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, err := a.ReportService.GetDiverReportData(r.Context(), diverAddr, repID)
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

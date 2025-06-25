package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// General information about the application
// Also grant a information of connected divers
func (a *App) Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := `Hi! My name is Michman and I can manage all of your databases in a single API.`

		divers, err := a.Service.GetReports(r.Context())
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(divers) > 0 {
			text += `<table style="border-style: solid;"><tr><th>name</th><th>addres</th><th>status</th></tr>`
			for _, v := range divers {
				status := "<a>&#128993;</a>"
				// if v.CheckConn() {
				// 	status = "<a>&#128994;</a>"
				// } else {
				// 	status = "<a>&#128308;</a>"
				// }
				text += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td></tr>", v.Name, status)
			}
			text += "</table>"
		}

		w.Header().Add("content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(text))
	}
}

// Get all divers from inmemory storage
func (a *App) GetReports() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		divers, err := a.Service.GetReports(r.Context())
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(divers); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

// Execute report's query to diver thrue MQ
func (a *App) ExecuteReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		repID := chi.URLParam(r, "id")

		data, err := a.Service.ExecuteReport(r.Context(), repID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "status bad request"})
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "status bad request"})
			return
		}
	}
}

func errorJSON(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
}

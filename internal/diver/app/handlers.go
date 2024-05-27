package app

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/imirjar/Michman/internal/diver/models"
)

func (a *App) ExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var report models.Report
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		log.Println("HANDLER ExecuteHandler Decode ERROR", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := a.service.Execute(context.Background(), report.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	report.Data = result
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(report); err != nil {
		log.Println("HANDLER ExecuteHandler Encode ERROR", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) ReportsListHandler(w http.ResponseWriter, r *http.Request) {

	result, err := a.service.ReportsList(r.Context())
	if err != nil {
		log.Println("HANDLER ExecuteHandler Encode ERROR", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/imirjar/Michman/internal/auth/models"
)

func (a *App) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, I'm auth app"))
}

func (a *App) CreateJWT(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := a.service.BuildJWTString(r.Context(), credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func (a *App) ValidateJWT(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("jwt")

	userID := a.service.GetUserID(r.Context(), token)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(userID)))
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(credentials)))
}

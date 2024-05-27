package app

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := a.service.BuildJWTString(r.Context(), credentials)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func (a *App) ValidateJWT(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("jwt")

	if err := a.service.VerifyToken(r.Context(), token); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SUCCESS"))
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprint(credentials)))
}

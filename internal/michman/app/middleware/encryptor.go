package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func Authenticator(secret, authPath string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if secret != "" {
				hashHeader := r.Header.Get("Authorization")

				if !auth(hashHeader, authPath) {
					http.Error(w, "Auth failed", http.StatusForbidden)
					return
				}

			}
			next.ServeHTTP(w, r)
		})
	}
}

func auth(header, authPath string) bool {
	absPath := fmt.Sprintf("http://" + authPath + "/token/validate/")
	client := http.Client{}

	req, err := http.NewRequest("POST", absPath, nil)
	if err != nil {
		return false
	}

	req.Header.Add("jwt", header)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		// log.Print(response.StatusCode)
		return true
	} else {
		log.Print(response.StatusCode)
		return false
	}
}

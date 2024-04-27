package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func Encryptor(secret, authPath string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if secret != "" {
				hashHeader := r.Header.Get("jwt")
				if hashHeader != "" {
					if err := auth(hashHeader, authPath); err != nil {
						http.Error(w, err.Error(), http.StatusForbidden)
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func auth(header, authPath string) error {
	absPath := fmt.Sprintf("http://" + authPath + "/token/view/")
	client := http.Client{}

	req, err := http.NewRequest("POST", absPath, nil)
	if err != nil {
		// log.Print("I CANT REQUEST")
		return err
	}

	req.Header.Add("jwt", header)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		return nil
	} else {
		return err
	}
}

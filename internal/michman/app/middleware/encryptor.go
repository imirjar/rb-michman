package middleware

import (
	"log"
	"net/http"
)

func Encryptor(secret, authPath string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if secret != "" {
				hashHeader := r.Header.Get("jwt")

				if hashHeader != "" {

					if !auth(hashHeader, "http//"+authPath+"/token/view/") {
						log.Print("YOU SHALL NOT PASS!")
						http.Error(w, "unvalid secret", http.StatusForbidden)
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func auth(header, authPath string) bool {
	client := http.Client{}
	req, err := http.NewRequest("POST", authPath, nil)
	req.Header.Add("jwt", header)
	if err != nil {
		log.Print("I CANT REQUEST")
		return false
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		return true
	} else {
		return false
	}
}
